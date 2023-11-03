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

	// New SOR
	newSOR, err := client.NewSOR(models.NewSORRequest{
		Symbol:      "BNBFDUSD",
		Side:        "BUY",
		Type:        "LIMIT",
		Price:       22.5,
		Quantity:    1,
		Timestamp:   time.Now().UnixMilli(),
		TimeInForce: "GTC",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newSOR)
}
