package environment

type Bryan struct {
}

type BryanData struct {
	PoPTokens []int
}

func (e *Bryan) GenerateData(phase uint64, totWallet int) interface{} {
	var PoPTokens []int
	for i := 0; i < totWallet; i += 2 {
		PoPTokens = append(PoPTokens, i)
	}
	return BryanData{PoPTokens}
}
