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

	// for the example apiKey and secretKey are empty
	client := v3.NewBinanceClient(apiKey, secretKey)

	err = client.NewOrderTest(models.OrderRequest{Symbol: "BTCUSDT"})

	if err != nil {
		fmt.Errorf(err.Error())
	}

}
