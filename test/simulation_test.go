package test

import (
	"../model"
	"../simulation"
	"log"
	"testing"
)

func Run(name string, interval int, totWallet int, random bool, t *testing.T) {

	log.Println("Starting " + name + " environment")

	// Initialization
	simulation.GenerateWallets(totWallet)
	currencyBalance := model.Factory("BalanceModel", "CNY")
	currencyUtxo := model.Factory("UtxoModel", "USD")

	// Starting the simulation
	simulation.Simulate(name, interval, currencyBalance, 0, random)
	simulation.Simulate(name, interval, currencyUtxo, 0, random)

	// Comparing the results in both models
	for i := 0; i < totWallet; i++ {
		valueBalance := currencyBalance.BalanceOf(simulation.GetWallet(i))
		valueUtxo := currencyUtxo.BalanceOf(simulation.GetWallet(i))
		log.Printf("Wallet balance of %s is %d in Balance Model and %d in UTXO model \n", simulation.GetWallet(i).Addr, valueBalance, valueUtxo)
		if valueBalance != valueUtxo {
			t.Errorf("Balance is not consistent in the two models.")
		}
	}
}

func TestSimulation(t *testing.T) {
	//log.SetOutput(ioutil.Discard)
	Run("bitcoin", 20, 5, false, t)
	Run("ethereum", 20, 5, false, t)
	Run("stake", 20, 5, true, t)
	Run("bryan", 20, 5, false, t)
}
