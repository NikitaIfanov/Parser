package main

import (
	"main.go/pkg/Exchange"
	"time"
)

func main() {
	pair := Exchange.EnterPair()
	mapAccount := MakeAccount()
	for {
		data := pair.Make()

		MakeMoney(pair, data, mapAccount)

		time.Sleep(1 * time.Second)

	}

}
