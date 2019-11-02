package model

type BalanceModel struct {
	Symbol   string
	Balances map[string]uint64
}

// Create currency from thin air
func (c *BalanceModel) Mint(u Wallet, amount uint64) {
	if c.Balances[u.Addr]+amount < c.Balances[u.Addr] {
		return
	}
	c.Balances[u.Addr] += amount
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
