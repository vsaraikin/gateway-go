package main

import (
	"fmt"
	"gateaway/binance/models"
	"gateaway/binance/ws"
	"os"
	"os/signal"
)

func main() {
	// Load config from ./config/.env
	//apiKey, secretKey, err := config.LoadEnv()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// for the example apiKey and secretKey are empty
	client := ws.NewBinanceWsClient("", "")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	dataHandler := func(e *models.DepthEvent) {
		fmt.Println(e)
	}

	err, done := client.SubscribeDepth("btcusdt", dataHandler)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	<-interrupt        // Interrupt by CTRL+C
	done <- struct{}{} // Graceful shutdown closing subscription
}
