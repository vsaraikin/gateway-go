package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
	"gateaway/config"
)

// TODO:
// 1. Change prices and quantity to `decimal.Decimal`
// 2. Add lawyers such as signedPost, signedGet, unsignedPost....
// 3. Move out executeRequest from Binance class
// 4. Measure time exec.
// 5. Timestamp should automatically be signed

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
