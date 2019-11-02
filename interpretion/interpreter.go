package interpretion

import (
	"encoding/json"
	"io/ioutil"
)

type Rewards struct {
	Times     string
	Formula   string
	Target    string
	//Authority string
	Condition string
}

type Issuance struct {
	Base      string
	Frequency int
	Rewards   []Rewards
}

func Interpret(path string) Issuance {
	file, _ := ioutil.ReadFile(path)
	res := Issuance{}
	_ = json.Unmarshal([]byte(file), &res)
	return res
}
