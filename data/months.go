package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type Day struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Month struct {
	Month string `json:"month"`
	Id    int    `json:"id"`
	Real  string `json:"real"`
	Days  []Day  `json:"days"`
}

//go:embed months.json
var monthsfile []byte

var Months []Month

func init() {
	Months = []Month{}
	if err := json.Unmarshal(monthsfile, &Months); err != nil {
		fmt.Println("Cannot load months json")
	}
}
