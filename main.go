package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
	"gateaway/config"
)

func main() {
	// Load config from ./config/.env
	apiKey, secretKey, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := v3.NewBinanceClient(apiKey, secretKey)

	depth, err := client.GetDepth(models.DepthRequest{
		Symbol: "SOLUSDT",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(depth)
}
