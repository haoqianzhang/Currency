package test

import (
	"../interpretion"
	"testing"
)

func TestInterpretBitcoin(t *testing.T) {
	interpretion.Interpret("../configuration/bitcoin.json")
	interpretion.Interpret("../configuration/ethereum.json")
	interpretion.Interpret("../configuration/stake.json")
	interpretion.Interpret("../configuration/bryan.json")
}
