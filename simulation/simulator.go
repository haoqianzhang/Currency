package simulation

import (
	"../interpretion"
	"../model"
	"./environment"
	"fmt"
	"github.com/antonmedv/expr"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

func GenerateWallets(totWallet int) {
	environment.GenerateWallets(totWallet)
}

func GetWallet(index int) model.Wallet {
	return environment.Wallets[index]
}

func GetAddr(index int) string {
	return GetWallet(index).Addr
}

type Simulator struct {
	Name string
	Interval int
	Currency model.Currency
	Seed int64
	Random bool
}

func (s *Simulator) Rebase(amount uint64) float64{
	return float64(amount)/math.Pow10(s.Currency.ValueOfDecimal())
}

func (s *Simulator) Run() {

	// Read the JSON configuration file
	res := interpretion.Interpret("../configuration/" + s.Name + ".json")

	// Init
	environment.InitSeed(s.Seed)
	env := environment.Factory(s.Name, s.Currency)

	// Randomly give each wallet some initial money
	if s.Random == true {
		for i := 0; i < len(environment.Wallets); i++ {
			s.Currency.Mint(environment.Wallets[i], uint64(rand.Intn(100 * int(math.Pow10(s.Currency.ValueOfDecimal())))))
		}
	}

	for phase := 0; phase < s.Interval; phase += res.Frequency {

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
				formula := strconv.FormatInt(int64(math.Pow10(s.Currency.ValueOfDecimal())),10)+ " * " + res.Rewards[r].Formula
				program, _ := expr.Compile(formula, expr.Env(table), expr.AsInt64())
				amount, _ := expr.Run(program, table)

				// Calculating reward target
				target, _ := expr.Eval(res.Rewards[r].Target, table)

				// Mint new coins
				log.Printf("Mint %.2f coins to %s based on formula: %s\n", s.Rebase(uint64(amount.(int64))), GetAddr(target.(int)),res.Rewards[r].Formula)
				s.Currency.Mint(environment.Wallets[target.(int)], uint64(amount.(int64)))

				// Generate a transfer transaction randomly
				transaction := environment.GenerateTransaction(len(environment.Wallets))
				from := transaction.From
				to := transaction.To
				value := transaction.Value
				log.Printf("Transfer %.2f coins from %s to %s\n", s.Rebase(value), GetAddr(from), GetAddr(to))
				s.Currency.Transfer(environment.Wallets[from], environment.Wallets[to], value)
			}
		}

		fmt.Printf("total coin %.2f at phase %d\n", s.Currency.TotalSupply(), phase)
	}
}
