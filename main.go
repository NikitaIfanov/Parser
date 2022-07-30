package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	Binance = 1
	Huobi   = 2
	Okex    = 3
	ByBit   = 4
)

func main() {
	p := EnterPair()
	var count int
	for {
		Data := make([]DataFloat, 0)

		Data = append(Data, p.GetBinance(), p.GetHuobi(), p.GetOkex(), p.GetByBit())
		fmt.Println(Data)
		MaxSale := Data[0]
		MinBuy := Data[0]
		for _, num := range Data {
			switch {
			case num.SalePrice > MaxSale.SalePrice:
				MaxSale = num
			case num.BuyPrice < MinBuy.BuyPrice:
				MinBuy = num
			}
		}
		fmt.Println(MaxSale, MinBuy)
		if MaxSale.SalePrice-MinBuy.BuyPrice > 20 {
			count++
			fmt.Println(count)
		}

		time.Sleep(1 * time.Second)

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

func JsonUnmarshal(body []byte, targets interface{}) {
	err := json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
}

func EnterPair() Pair {
	var FirstExchange, SecondExchange string
	fmt.Println("Enter a Cryptocurrency Pair Example:'BTC USDT'")
	fmt.Scanf("%s %s", &FirstExchange, &SecondExchange)
	return Pair{
		Binance: fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Huobi:   fmt.Sprintf("%s%s", strings.ToLower(FirstExchange), strings.ToLower(SecondExchange)),
		Okex:    fmt.Sprintf("%s-%s", FirstExchange, SecondExchange),
		ByBit:   fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
	}

}

func ParseFloat(s string) float64 {
	data, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type DataFloat struct {
	Flag      int
	BuyPrice  float64
	SalePrice float64
}

type Pair struct {
	Binance string
	Huobi   string
	Okex    string
	ByBit   string
}
