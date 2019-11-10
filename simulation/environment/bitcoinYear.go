package environment

type BitcoinYear struct {
}

type BitcoinYearData struct {
	Year          uint64
	BitcoinSupply int
}

func (e *BitcoinYear) GenerateData(phase uint64, totWallet int) interface{} {

	return BitcoinYearData{phase, 0}
}
