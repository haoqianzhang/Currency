package environment

import (
	"../../model"
)

type Stake struct {
	Currency model.Currency
}

type Stakeholder struct {
	Amount  uint64
	Address int
}

type StakeData struct {
	Stakeholders []Stakeholder
}

func (e *Stake) GenerateData(phase uint64, totWallet int) interface{} {
	var Stakeholders []Stakeholder
	for i := 0; i < totWallet; i++ {
		Stakeholders = append(Stakeholders, Stakeholder{e.Currency.BalanceOf(Wallets[i]), i})
	}
	return StakeData{Stakeholders}
}
