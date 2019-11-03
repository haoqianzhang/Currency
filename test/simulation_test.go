package test

import (
	"github.com/haoqianzhang/currency/model"
	"github.com/haoqianzhang/currency/simulation"
	"io/ioutil"
	"log"
	"testing"
)

func Run(name string, interval int, totWallet int, random bool, decimal int, t *testing.T) {

	log.Println("Starting " + name + " environment")

	// Initialization
	simulation.GenerateWallets(totWallet)

	currencyBalance := model.Factory("BalanceModel", "CNY", decimal)
	simBalance := simulation.Simulator{Name: name, Interval: interval, Currency: currencyBalance, Seed: 0, Random: random}

	currencyUtxo := model.Factory("UtxoModel", "USD", decimal)
	simUtxo := simulation.Simulator{Name: name, Interval: interval, Currency: currencyUtxo, Seed: 0, Random: random}

	simBalance.Run()
	simUtxo.Run()

	// Comparing the results in both models
	for i := 0; i < totWallet; i++ {
		valueBalance := currencyBalance.BalanceOf(simulation.GetWallet(i))
		valueUtxo := currencyUtxo.BalanceOf(simulation.GetWallet(i))
		valueRebasedBalance := simBalance.Rebase(valueBalance)
		valueRebasedUtxo := simUtxo.Rebase(valueUtxo)
		log.Printf("Wallet balance of %s is %.2f in Balance Model and %.2f in UTXO model \n", simulation.GetAddr(i), valueRebasedBalance, valueRebasedUtxo)
		if valueBalance != valueUtxo {
			t.Errorf("Balance is not consistent in the two models.")
		}
	}
}

func TestSimulation(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	Run("bitcoin", 20, 5, false,8, t)
	Run("bitcoinYear", 20, 1, false,8, t)
	Run("ethereum", 20, 5, false,8, t)
	Run("stake", 20, 5, true,8, t)
	Run("bryan", 20, 5, false, 8, t)
}
