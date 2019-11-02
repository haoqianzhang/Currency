package simulation

import (
	"../interpretion"
	"../model"
	"./environment"
	"github.com/antonmedv/expr"
	"log"
	"math/rand"
	"reflect"
)

func GenerateWallets(totWallet int) {
	environment.GenerateWallets(totWallet)
}

func GetWallet(index int) model.Wallet {
	return environment.Wallets[index]
}

func Simulate(name string, interval int, currency model.Currency, seed int64, random bool) {

	// Read the JSON configuration file
	res := interpretion.Interpret("../configuration/" + name + ".json")

	// Init
	environment.InitSeed(seed)
	env := environment.Factory(name, currency)

	// Randomly give each wallet some initial money
	if random == true {
		for i := 0; i < len(environment.Wallets); i++ {
			currency.Mint(environment.Wallets[i], uint64(rand.Intn(100)))
		}
	}

	for phase := 0; phase < interval; phase += res.Frequency {

		data := env.GenerateData(uint64(phase), len(environment.Wallets))

		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)

		table := map[string]interface{}{}

		for k := 0; k < t.NumField(); k++ {
			table[t.Field(k).Name] = v.Field(k).Interface()
		}

		for r := 0; r < len(res.Rewards); r++ {
			var input string

			// Checking condition
			input = "true"
			if len(res.Rewards[r].Condition) > 0 {
				input = res.Rewards[r].Condition
			}
			condition, _ := expr.Eval(input, table)
			if condition == false {
				continue
			}

			// Calculating times
			input = "1"
			if len(res.Rewards[r].Times) > 0 {
				input = res.Rewards[r].Times
			}
			program, _ := expr.Compile(input, expr.Env(table), expr.AsInt64())
			times, _ := expr.Run(program, table)

			for i := 0; i < int(times.(int64)); i++ {

				table["i"] = i

				// Calculating reward formula
				program, _ := expr.Compile(res.Rewards[r].Formula, expr.Env(table), expr.AsInt64())
				amount, _ := expr.Run(program, table)

				//// Calculating reward target
				target, _ := expr.Eval(res.Rewards[r].Target, table)

				// Mint new coins
				log.Println("Mint", amount, "coins to"+environment.Wallets[target.(int)].Addr, "based on formula:", res.Rewards[r].Formula)
				currency.Mint(environment.Wallets[target.(int)], uint64(amount.(int64)))

				// Generate a transfer transaction randomly
				transaction := environment.GenerateTransaction(len(environment.Wallets))
				from := transaction.From
				to := transaction.To
				value := transaction.Value
				log.Println("Transfer", value, "coins from "+environment.Wallets[from].Addr, "to", environment.Wallets[to].Addr)
				currency.Transfer(environment.Wallets[from], environment.Wallets[to], value)
			}
		}
	}
}
