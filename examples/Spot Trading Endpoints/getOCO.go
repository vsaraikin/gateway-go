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

	// New OCO
	stopLimit := 22.5
	newOCO, err := client.NewOCO(models.NewOCORequest{
		Symbol:               "SOLUSDT",
		Side:                 "BUY",
		Price:                20,
		Quantity:             1,
		StopPrice:            40,
		Timestamp:            time.Now().UnixMilli(),
		StopLimitPrice:       &stopLimit,
		StopLimitTimeInForce: "GTC",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(newOCO)
	time.Sleep(2 * time.Second)

	// Get OCO
	canceledOrderList, err := client.GetOCO(models.GetOCORequest{
		Timestamp:   time.Now().UnixMilli(),
		OrderListID: &newOCO.OrderListId,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(canceledOrderList)

}
