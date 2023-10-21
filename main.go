package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
	"gateaway/config"
	"time"
)

// TODO:
// 1. Change prices and quantity to `decimal.Decimal`
// 2. Add lawyers such as signedPost, signedGet, unsignedPost....
// 3. Move out executeRequest from Binance class

func main() {
	// Load config from ./config/.env
	apiKey, secretKey, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	// for the example apiKey and secretKey are empty
	client := v3.NewBinanceClient(apiKey, secretKey)

	order, err := client.NewOrder(models.OrderRequest{
		Symbol:     "BTCUSDT",
		Side:       "BUY",
		Type:       "MARKET",
		Quantity:   1,
		RecvWindow: 10000,
		Timestamp:  time.Now().UnixMilli()})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(order)
}
