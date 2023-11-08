package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
)

func main() {

	client := v3.NewBinanceClient("", "")

	trades, err := client.HistoricalTrades(models.TradesRequest{
		Symbol: "SOLUSDT",
		Limit:  3,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(trades)

}
