package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
	"gateaway/config"
	"time"
)

func main() {
	// Load config from ./config/.env
	apiKey, secretKey, err := config.LoadEnv()
	if err != nil {

		fmt.Println(err)
		return
	}

	// for the example apiKey and secretKey are empty
	client := v3.NewBinanceClient(apiKey, secretKey)

	err = client.NewOrderTest(models.OrderRequest{
		Symbol:     "BTCUSDT",
		Side:       models.BUY,
		Type:       models.MARKET,
		Quantity:   1,
		RecvWindow: 10000,
		Timestamp:  time.Now().UnixMilli()})

	if err != nil {
		fmt.Println(err.Error())
	}

	info, err := client.GetExchangeInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)

}
