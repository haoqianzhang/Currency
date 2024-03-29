package model

import "math"

type BalanceModel struct {
	Symbol   string
	Balances map[string]uint64
	Decimal  int
	Supply   uint64
}

// Create currency from thin air
func (c *BalanceModel) Mint(u Wallet, amount uint64) {
	if c.Balances[u.Addr]+amount < c.Balances[u.Addr] {
		return
	}
	c.Balances[u.Addr] += amount
	c.Supply += amount
}

func (c *BalanceModel) Transfer(from Wallet, to Wallet, value uint64) bool {
	if c.Balances[from.Addr] < value {
		return false
	}
	c.Balances[from.Addr] -= value
	c.Balances[to.Addr] += value
	return true
}

func (c *BalanceModel) BalanceOf(u Wallet) uint64 {
	return c.Balances[u.Addr]
}

func (c *BalanceModel) RebasedBalanceOf(u Wallet) float64 {
	return float64(c.BalanceOf(u)) / math.Pow10(c.Decimal)
}

func (c *BalanceModel) TotalSupply() float64 {
	return float64(c.Supply) / math.Pow10(c.Decimal)
}

func (c *BalanceModel) ValueOfDecimal() int {
	return c.Decimal
}

func (c *BalanceModel) SetInitialSupply(initialSupply uint64) {
	c.Supply = initialSupply * uint64(math.Pow10(c.Decimal))
}
