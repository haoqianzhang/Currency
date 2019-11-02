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
	//var parameters []map[string]interface{}
	//for i := 0; i < interval; i++ {
	//	var m = map[string]interface{}{}
	//	m["Height"] = uint64(i)
	//	m["BlockMiner"] = rand.Intn(walletsNum)
	//	parameters = append(parameters, m)
	//	//parameters[i]["Height"] = uint64(i)
	//	//parameters[i]["BlockMiner"] = rand.Intn(walletsNum)
	//	//blocks = append(blocks, Parameter{uint64(i), rand.Intn(walletsNum)})
	//}
	//return parameters
}
