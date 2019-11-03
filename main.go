package main

import (
	"encoding/csv"
	"fmt"
	"github.com/haoqianzhang/currency/model"
	"github.com/haoqianzhang/currency/simulation"
	"log"
	"os"
	"strconv"
)

func Run(name string, interval int, totWallet int, random bool, decimal int) {

	log.Println("Starting " + name + " environment")

	// Initialization
	simulation.GenerateWallets(totWallet)
	currencyBalance := model.Factory("BalanceModel", "CNY", decimal)

	// Starting the simulation
	simBalance := simulation.Simulator{Name: name, Interval: interval, Currency: currencyBalance, Seed: 0, Random: random}
	simBalance.Run()

	f, err := os.Create(name + ".csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// UTF 8
	_, _ = f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	var data [][]string
	for i :=0 ; i<len(simBalance.Supply) ; i++ {
		appendData := []string{strconv.Itoa(i),fmt.Sprintf("%.2f", simBalance.Supply[i])}
		if i>0{
			inflation := (simBalance.Supply[i] - simBalance.Supply[i-1]) / simBalance.Supply[i-1]
			appendData = append(appendData, fmt.Sprintf("%.4f", inflation * 100.0))
		}
		data = append(data, appendData)
	}
	_ = w.WriteAll(data)
	w.Flush()
}

func main() {
	Run("bitcoinYear", 100, 1, false,8)
}