package main

import (
	"fmt"
	"main.go/pkg/Exchange"
)

type Account struct {
	Flag          int
	BalanceUSDT   float64
	BalanceCrypto float64
	SaleAllowed   bool
	BuyAllowed    bool
}

func (account *Account) Buy(data Exchange.DataFloat) {

	percent := data.BuyPrice / 1000
	account.BalanceUSDT = account.BalanceUSDT - 0.1*data.BuyPrice - percent
	account.BalanceCrypto += 0.1
	account.SaleAllowed = true
	account.BuyAllowed = false

}
func (account *Account) Sale(data Exchange.DataFloat) {

	percent := data.SalePrice / 1000
	account.BalanceUSDT = account.BalanceUSDT + 0.1*data.SalePrice - percent
	account.BalanceCrypto -= 0.1
	account.SaleAllowed = false
	account.BuyAllowed = true

}

func MakeMoney(pair Exchange.Pair, Data []Exchange.DataFloat, mapAccount map[int]*Account) {
	MaxSale := Data[0]
	MinBuy := Data[0]
	for _, num := range Data {
		switch {
		case num.SalePrice > MaxSale.SalePrice:
			MaxSale = num

		case num.BuyPrice < MinBuy.BuyPrice:
			MinBuy = num

		}
		if MaxSale.SalePrice-MinBuy.BuyPrice > pair.Difference {
			if mapAccount[MaxSale.Flag].SaleAllowed == true && mapAccount[MinBuy.Flag].BuyAllowed == true {
				mapAccount[MaxSale.Flag].Sale(MaxSale)
				mapAccount[MinBuy.Flag].Buy(MinBuy)
				fmt.Println(mapAccount[1], mapAccount[2], mapAccount[3], mapAccount[4], mapAccount[5])
			}
		}
	}

}

func MakeAccount() map[int]*Account {
	a := make(map[int]*Account)
	//2500$ and 0.6 ETH
	a[1] = &Account{
		Flag:          Exchange.Binance,
		BalanceUSDT:   500,
		BalanceCrypto: 0.12,
		SaleAllowed:   true,
		BuyAllowed:    true,
	}
	a[2] = &Account{
		Flag:          Exchange.Huobi,
		BalanceUSDT:   500,
		BalanceCrypto: 0.12,
		SaleAllowed:   true,
		BuyAllowed:    true,
	}
	a[3] = &Account{
		Flag:          Exchange.Okex,
		BalanceUSDT:   500,
		BalanceCrypto: 0.12,
		SaleAllowed:   true,
		BuyAllowed:    true,
	}
	a[4] = &Account{
		Flag:          Exchange.ByBit,
		BalanceUSDT:   500,
		BalanceCrypto: 0.12,
		SaleAllowed:   true,
		BuyAllowed:    true,
	}
	a[5] = &Account{
		Flag:          Exchange.Kraken,
		BalanceUSDT:   500,
		BalanceCrypto: 0.12,
		SaleAllowed:   true,
		BuyAllowed:    true,
	}
	return a
}
