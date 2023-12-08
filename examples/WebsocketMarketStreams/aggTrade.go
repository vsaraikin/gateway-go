package main

import (
	"fmt"
	"gateaway/binance/ws"
	"gateaway/binance/ws/models"
	"os"
	"os/signal"
)

func main() {
	// Endpoint does not require auth
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
