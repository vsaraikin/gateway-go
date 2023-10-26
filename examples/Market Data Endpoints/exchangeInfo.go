package main

import (
	"fmt"
	v3 "gateaway/binance/v3"
)

func main() {
	// Endpoint does not require auth
	client := v3.NewBinanceClient("", "")

	response, err := client.GetExchangeInfo()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(response)
}
