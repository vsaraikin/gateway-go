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

	client := v3.NewBinanceClient(apiKey, secretKey)

	order, err := client.NewOrderTest(models.OrderRequest{
		Symbol:     "ETHUSDT",
		Side:       "BUY",
		Type:       "MARKET",
		Quantity:   0.1,
		RecvWindow: 10000,
		Timestamp:  time.Now().UnixMilli()})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(order)
}
