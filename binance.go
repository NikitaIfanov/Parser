package main

func (p Pair) GetBinance() DataFloat {
	body := GetJson("https://api.binance.com/api/v3/ticker/bookTicker")

	targetsBinance := []BinanceJson{}

	JsonUnmarshal(body, &targetsBinance)
	binance := DataFloat{Flag: Binance}
	for _, t := range targetsBinance {
		if t.Symbol == p.Binance {
			binance.BuyPrice = ParseFloat(t.BuyPrice)
			binance.SalePrice = ParseFloat(t.SalePrice)
			break
		}

	}
	return binance
}

type BinanceJson struct {
	Symbol    string `json:"symbol"`
	BuyPrice  string `json:"bidPrice"`
	SalePrice string `json:"askPrice"`
}
