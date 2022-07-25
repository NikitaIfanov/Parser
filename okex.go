package main

import (
	"fmt"
)

func (p Pair) GetOkex() {
	body := GetJson("https://www.okx.com/api/v5/market/tickers/?instType=SPOT")

	targets := OkexJson{}

	JsonUnmarshal(body, &targets)

	for _, t := range targets.Data {
		if t.Symbol == p.Okex {
			fmt.Println("Okex" + " " + t.Symbol + " = " + t.Price)
			break
		}
	}
}

type OkexJson struct {
	Data []struct {
		Symbol string `json:"instId"`
		Price  string `json:"askPx"`
	} `json:"data"`
}
