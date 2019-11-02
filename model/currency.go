package model

type Currency interface {
	Mint(u Wallet, amount uint64)
	BalanceOf(u Wallet) uint64
	Transfer(from Wallet, to Wallet, value uint64) bool
	TotalSupply() float64
	ValueOfDecimal() int
	RebasedBalanceOf(u Wallet) float64
}

func Factory(name string, symbol string, decimal int) Currency {
	switch name {
	case "BalanceModel":
		return &BalanceModel{symbol, make(map[string]uint64), decimal,0}
	case "UtxoModel":
		return &UtxoModel{symbol, make(map[string][]*Output), []Transaction{}, decimal,0}
	default:
		panic("No such model")
	}
}