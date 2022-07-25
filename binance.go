package main

import (
	"fmt"
)

func (p Pair) GetBinance() {
	body := GetJson("https://api.binance.com/api/v3/ticker/price")

	targets := []BinanceJson{}

	JsonUnmarshal(body, &targets)

	for _, t := range targets {
		if t.Symbol == p.Binance {
			fmt.Println("Binance" + " " + t.Symbol + " = " + t.Price)
			break
		}
	}
}

type BinanceJson struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
