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

	// All OCO list
	canceledOrderList, err := client.AllOCOList(models.AllOCOListRequest{
		Timestamp: time.Now().UnixMilli(),
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(canceledOrderList)

}
