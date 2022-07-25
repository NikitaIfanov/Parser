package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	p := EnterPair()

	for {

		p.GetBinance()
		p.GetHuobi()
		p.GetOkex()

		time.Sleep(10 * time.Second)

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

type Pair struct {
	Binance string
	Huobi   string
	Okex    string
}

func EnterPair() Pair {
	var FirstExchange, SecondExchange string
	fmt.Println("Enter a Cryptocurrency Pair Example:'BTC USDT'")
	fmt.Scanf("%s %s", &FirstExchange, &SecondExchange)
	return Pair{
		Binance: fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Huobi:   fmt.Sprintf("%s%s", strings.ToLower(FirstExchange), strings.ToLower(SecondExchange)),
		Okex:    fmt.Sprintf("%s-%s", FirstExchange, SecondExchange),
	}

}
