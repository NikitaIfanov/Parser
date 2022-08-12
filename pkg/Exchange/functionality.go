package Exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	Binance = 1
	Huobi   = 2
	Okex    = 3
	ByBit   = 4
	Kraken  = 5
)

type DataFloat struct {
	Flag      int
	BuyPrice  float64
	SalePrice float64
}

type Pair struct {
	Difference float64
	Binance    string
	Huobi      string
	Okex       string
	ByBit      string
	Kraken     string
}

func JsonUnmarshal(body []byte, targets interface{}) {
	err := json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
}

func GetJson(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return body
}

func (p Pair) Make() []DataFloat {
	Data := make([]DataFloat, 0, 5)

	Data = append(Data,

		p.GetBinance(), p.GetHuobi(), p.GetOkex(), p.GetByBit(), p.GetKraken(),
	)
	return Data
}

func EnterPair() Pair {
	var (
		FirstExchange, SecondExchange string
		Difference                    float64
	)

	fmt.Println("Enter a Cryptocurrency Pair & Difference Example:'BTC USDT 10'")
	fmt.Scanf("%s %s %f", &FirstExchange, &SecondExchange, &Difference)
	return Pair{
		Difference: Difference,
		Binance:    fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Huobi:      fmt.Sprintf("%s%s", strings.ToLower(FirstExchange), strings.ToLower(SecondExchange)),
		Okex:       fmt.Sprintf("%s-%s", FirstExchange, SecondExchange),
		ByBit:      fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Kraken:     fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
	}

}

func ParseFloat(s string) float64 {
	data, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
