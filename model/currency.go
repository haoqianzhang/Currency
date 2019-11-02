package model

type Currency interface {
	Mint(u Wallet, amount uint64)
	BalanceOf(u Wallet) uint64
	Transfer(from Wallet, to Wallet, value uint64) bool
}

func Factory(name string, symbol string) Currency {
	switch name {
	case "BalanceModel":
		return &BalanceModel{symbol, make(map[string]uint64)}
	case "UtxoModel":
		return &UtxoModel{symbol, make(map[string][]*Output), []Transaction{}}
	default:
		panic("No such model")
	}
}
