package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
)

func main() {

	client := v3.NewBinanceClient("", "")

	trades, err := client.Klines(models.KlineRequest{
		Symbol:   "SOLUSDT",
		Interval: "1m",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(trades)

}
