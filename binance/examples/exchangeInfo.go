package main

import (
	"fmt"
	v3 "gateaway/binance/v3"
)

func main() {
	// Load config from ./config/.env
	//apiKey, secretKey, err := config.LoadEnv()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// for the example apiKey and secretKey are empty
	client := v3.NewBinanceClient("", "")

	response, err := client.GetExchangeInfo()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(response)
}
