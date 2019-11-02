package model

import "math"

type UtxoModel struct {
	Symbol string
	UTXO   map[string][]*Output
	Ledger []Transaction
	Decimal int
	Supply uint64
}

type Output struct {
	Value uint64
	Addr  string
}

type Transaction struct {
	Inputs  []*Output
	Outputs []Output
}

// Create currency from thin air
func (c *UtxoModel) Mint(u Wallet, amount uint64) {
	var outputs = []Output{{amount, u.Addr}}
	c.Ledger = append(c.Ledger, Transaction{[]*Output{}, outputs})
	c.UTXO[u.Addr] = append(c.UTXO[u.Addr], &c.Ledger[len(c.Ledger)-1].Outputs[0])
	c.Supply += amount
}

func (c *UtxoModel) BalanceOf(u Wallet) uint64 {
	var balance uint64 = 0
	for i := 0; i < len(c.UTXO[u.Addr]); i++ {
		balance += c.UTXO[u.Addr][i].Value
	}
	return balance
}

func (c *UtxoModel) Transfer(from Wallet, to Wallet, value uint64) bool {
	// Not allow to transfer 0 coin
	if value == 0 {
		return false
	}

	// Step 1: checking if there is enough coin in owner's account
	if c.BalanceOf(from) < value {
		return false
	}

	// Step 2: creating new transaction and adding it to the Ledger
	var sum uint64 = 0
	var inputs []*Output
	var pos = 0
	for i := 0; i < len(c.UTXO[from.Addr]); i++ {
		sum += c.UTXO[from.Addr][i].Value
		inputs = append(inputs, c.UTXO[from.Addr][i])
		if sum >= value {
			pos = i
			break
		}
	}
	var outputs = []Output{{value, to.Addr}}
	if sum > value {
		outputs = append(outputs, Output{sum - value, from.Addr})
	}
	var transaction = Transaction{inputs, outputs}
	c.Ledger = append(c.Ledger, transaction)

	// Step 3: update users' UTXO
	c.UTXO[from.Addr] = c.UTXO[from.Addr][pos+1:]
	c.UTXO[to.Addr] = append(c.UTXO[to.Addr], &c.Ledger[len(c.Ledger)-1].Outputs[0])
	if sum > value {
		c.UTXO[from.Addr] = append(c.UTXO[from.Addr], &c.Ledger[len(c.Ledger)-1].Outputs[1])
	}

	return true
}

func (c *UtxoModel) RebasedBalanceOf(u Wallet) float64 {
	return float64(c.BalanceOf(u))/math.Pow10(c.Decimal)
}

func (c *UtxoModel) ValueOfDecimal() int {
	return c.Decimal
}

func (c *UtxoModel) TotalSupply() float64 {
	return float64(c.Supply)/math.Pow10(c.Decimal)
}
