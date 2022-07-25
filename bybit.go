package main

import "fmt"

func (p Pair) GetByBit() {

	body := GetJson("https://api.bybit.com/v2/public/tickers")

	targets := ByBitJson{}

	JsonUnmarshal(body, &targets)

	for _, t := range targets.Data {
		if t.Symbol == p.ByBit {
			fmt.Println("ByBit" + " " + t.Symbol + " = " + t.Price)
			break
		}
	}
}

type ByBitJson struct {
	Data []struct {
		Symbol string `json:"symbol"`
		Price  string `json:"ask_price"`
	} `json:"result"`
}
