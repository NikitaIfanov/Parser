package main

func (p Pair) GetByBit() DataFloat {

	body := GetJson("https://api.bybit.com/v2/public/tickers")

	targetsByBit := ByBitJson{}

	JsonUnmarshal(body, &targetsByBit)
	byBit := DataFloat{Flag: ByBit}
	for _, t := range targetsByBit.Data {
		if t.Symbol == p.ByBit {
			byBit.BuyPrice = ParseFloat(t.BuyPrice)
			byBit.SalePrice = ParseFloat(t.SalePrice)
			break

		}
	}
	return byBit
}

type ByBitJson struct {
	Data []struct {
		Symbol    string `json:"symbol"`
		BuyPrice  string `json:"bid_price"`
		SalePrice string `json:"ask_price"`
	} `json:"result"`
}
