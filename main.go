package main

import (
	"fmt"
	"gateaway/binance/models"
	v3 "gateaway/binance/v3"
	"gateaway/config"
	"strconv"
	"time"
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

	err = client.NewOrderTest(models.OrderRequest{Symbol: "BTCUSDT", Side: v3.BUY, Type: v3.LIMIT, Quantity: fmt.Sprintf("%v", 1), RecvWindow: "10000", Timestamp: strconv.Itoa(int(time.Now().UnixMilli()))})

	if err != nil {
		fmt.Println(err.Error())
	}

}
