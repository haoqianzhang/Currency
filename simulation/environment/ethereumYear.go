package environment

type EthereumYear struct {
}

type EthereumYearData struct {
	BaseReward          int
	DailyUncleBlocks    int
	DailyBlocks         int
	UncleBlocksDistance float64
	EthSupply           int
}

func (e *EthereumYear) GenerateData(phase uint64, totWallet int) interface{} {
	return EthereumYearData{2, 450, 6400, 1.6, 0}
}
