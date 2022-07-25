package main

import (
	"fmt"
)

func (p Pair)GetBinance() {
	body := GetJson("https://api.binance.com/api/v3/ticker/price")

	targets := []MyJsonBin{}

	JsonUnmarshal(body, &targets)

	for _, t := range targets {
		if t.Symbol == p.Binance {
			fmt.Println("Binance" + " " + t.Symbol + " = " + t.Price)
			break
		}
	}
}

type MyJsonBin struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
