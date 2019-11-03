package test

import (
	"github.com/haoqianzhang/currency/interpretion"
	"testing"
)

func TestInterpretBitcoin(t *testing.T) {
	interpretion.Interpret("../configuration/bitcoin.json")
	interpretion.Interpret("../configuration/ethereum.json")
	interpretion.Interpret("../configuration/stake.json")
	interpretion.Interpret("../configuration/bryan.json")
}
