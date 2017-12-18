package main

import (
	"encoding/json"
	"vcm/api"
)

type Balance struct {
	Success bool   `json:"success"`
	BTC     string `json:"btc"`
	XRP     string `json:"xrp"`
	XEM     string `json:"xem"`
	ETH     string `json:"eth"`
	ETC     string `json:"etc"`
	LTC     string `json:"ltc"`
	FCT     string `json:"fct"`
	XMR     string `json:"xmr"`
	REP     string `json:"rep"`
	ZEC     string `json:"zec"`
	DASH    string `json:"dash"`
	BCH     string `json:"bch"`
}

func (b *Balance) Update() {
	result, err := api.FetchBalance()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(result, &b)
	if err != nil {
		panic(err)
	}
}
