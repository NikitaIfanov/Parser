package main

func (p Pair) GetOkex() DataFloat {
	body := GetJson("https://www.okx.com/api/v5/market/tickers/?instType=SPOT")

	targetsOkex := OkexJson{}

	JsonUnmarshal(body, &targetsOkex)
	okex := DataFloat{Flag: Okex}
	for _, t := range targetsOkex.Data {
		if t.Symbol == p.Okex {
			okex.BuyPrice = ParseFloat(t.BuyPrice)
			okex.SalePrice = ParseFloat(t.SalePrice)
			break
		}
	}
	return okex
}

type OkexJson struct {
	Data []struct {
		Symbol    string `json:"instId"`
		BuyPrice  string `json:"bidPx"`
		SalePrice string `json:"askPx"`
	}
}
