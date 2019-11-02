package environment

import (
	"math/rand"
)

type Ethereum struct {
}

type UncleBlock struct {
	Distance int
	Miner    int
}

type EthereumData struct {
	BaseReward  int
	BlockMiner  int
	UncleBlocks []UncleBlock
}

func (e *Ethereum) GenerateData(phase uint64, totWallet int) interface{} {
	var uncleBlocks []UncleBlock
	totUncleBlock := rand.Intn(3)
	for i := 0; i < totUncleBlock; i++ {
		uncleBlocks = append(uncleBlocks, UncleBlock{rand.Intn(6) + 2, rand.Intn(totWallet)})
	}
	return EthereumData{50, rand.Intn(totWallet), uncleBlocks}
}
