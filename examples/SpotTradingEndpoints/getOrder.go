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

	// Create limit order
	order, err := client.NewOrder(models.OrderRequest{
		Symbol:      "SOLUSDT",
		Side:        "BUY",
		Type:        "LIMIT",
		Price:       20,
		Quantity:    1,
		RecvWindow:  10000,
		Timestamp:   time.Now().UnixMilli(),
		TimeInForce: "GTC",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(order)

	time.Sleep(10 * time.Second)

	// Get order
	orderInfo, err := client.GetOrder(models.GetOrderRequest{
		Symbol:    "SOLUSDT",
		OrderID:   order.OrderId,
		Timestamp: time.Now().UnixMilli(),
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(orderInfo)
}
