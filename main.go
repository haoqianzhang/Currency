package main

import (
	"./model"
	"./simulation"
	"log"
)

func main() {
	//Run("bitcoin", 20, 5, false,8)
}

func Run(name string, interval int, totWallet int, random bool, decimal int) {

	log.Println("Starting " + name + " environment")

	// Initialization
	simulation.GenerateWallets(totWallet)
	currencyBalance := model.Factory("BalanceModel", "CNY", decimal)

	// Starting the simulation
	simBalance := simulation.Simulator{Name: name, Interval: interval, Currency: currencyBalance, Seed: 0, Random: random}
	simBalance.Run()
}
