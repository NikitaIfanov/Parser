package main

import (
	"fmt"
)

func (p Pair) GetHuobi() {
	body := GetJson("https://api.huobi.pro/market/tickers")

	targets := HuobiJson{}

	JsonUnmarshal(body, &targets)

	for _, t := range targets.Data {
		if t.Symbol == p.Huobi {
			fmt.Println("Huobi" + " " + t.Symbol + " = " + fmt.Sprintf("%f", t.Price))
			break
		}
	}
}

type HuobiJson struct {
	Data []struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"ask"`
	}
}
