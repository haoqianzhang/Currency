package environment

import "math/rand"

type Bitcoin struct {
}

type BitcoinData struct {
	Height     uint64
	BlockMiner int
}

func (e *Bitcoin) GenerateData(phase uint64, totWallet int) interface{} {

	return BitcoinData{phase, rand.Intn(totWallet)}
}
