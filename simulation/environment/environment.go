package environment

import (
	"../../model"
	"fmt"
	"math/rand"
)

type Transaction struct {
	From  int
	To    int
	Value uint64
}

var Wallets []model.Wallet

func GenerateWallets(totWallet int) {
	for i := 0; i < totWallet; i++ {
		addr := "wallet" + fmt.Sprintf("%08d", i)
		Wallets = append(Wallets, model.Wallet{addr})
	}
}

type Environment interface {
	GenerateData(phase uint64, totWallet int) interface{}
}

func InitSeed(seed int64) {
	rand.Seed(seed)
}

func GenerateTransaction(totWallet int) Transaction {
	return Transaction{rand.Intn(totWallet), rand.Intn(totWallet), uint64(rand.Intn(1000000000) + 1)}
}

func Factory(name string, currency model.Currency) Environment {
	switch name {
	case "bitcoin":
		return &Bitcoin{}
	case "ethereum":
		return &Ethereum{}
	case "stake":
		return &Stake{currency}
	case "bryan":
		return &Bryan{}
	default:
		panic("No such environment")
	}
}
