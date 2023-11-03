package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
)

func main() {
	client := v3.NewBinanceClient("", "")

	depth, err := client.GetDepth(models.DepthRequest{
		Symbol: "SOLUSDT",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(depth)
}
