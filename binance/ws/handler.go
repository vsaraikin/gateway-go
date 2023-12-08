package ws

import (
	"encoding/json"
	"fmt"
	"gateaway/binance/ws/models"
	"github.com/gorilla/websocket"
	"log"
)

type BinanceWsClient struct {
	APIKey        string
	Secret        string
	baseURL       string
	connections   *websocket.Conn
	subscriptions map[string]*websocket.Conn
	close         chan struct{} // channel to close connection
}

func NewBinanceWsClient(apiKey, secretKey string) *BinanceWsClient {
	return &BinanceWsClient{
		baseURL: "wss://stream.binance.com:9443/ws/",
		APIKey:  apiKey,
		Secret:  secretKey,
	}
}

func (c *BinanceWsClient) subscribe(url string, handler func(message []byte) error) (error, chan<- struct{}) {
	done := make(chan struct{})

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err, nil
	}

	go func() {
		// Wait to close connection in a separate goroutine while getting updates from WS
		defer close(done)  // Ensure done channel gets closed
		defer conn.Close() // Ensure connection gets closed

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				return
			}
			if err := handler(message); err != nil {
				log.Println("WebSocket handler error:", err)
				return
			}
		}
	}()

	return nil, done
}

type depthHandler func(e *models.DepthEvent)

func (c *BinanceWsClient) subscribeDepth(url string, handler depthHandler) (error, chan<- struct{}) {
	wsHandler := func(event []byte) error {
		depthEventRaw := new(models.DepthEventRaw)
		if err := json.Unmarshal(event, depthEventRaw); err != nil {
			log.Println("Error during JSON parsing:", err)
			return err
		}
		depthEvent := depthEventRaw.Transform()
		handler(depthEvent)
		return nil
	}
	return c.subscribe(url, wsHandler)
}

func (c *BinanceWsClient) SubscribeDepth(symbol string, handler depthHandler) (error, chan<- struct{}) {
	url := fmt.Sprintf("%s%s%s", c.baseURL, symbol, depth)
	return c.subscribeDepth(url, handler)
}

type aggTradeHandler func(e *models.AggTrade)

func (c *BinanceWsClient) subscribeAggTrade(url string, handler aggTradeHandler) (error, chan<- struct{}) {
	wsHandler := func(event []byte) error {
		aggTradesEvent := new(models.AggTrade)
		if err := json.Unmarshal(event, aggTradesEvent); err != nil {
			log.Println("Error during JSON parsing:", err)
			return err
		}
		handler(aggTradesEvent)
		return nil
	}
	return c.subscribe(url, wsHandler)
}

func (c *BinanceWsClient) SubscribeAggTrade(symbol string, handler aggTradeHandler) (error, chan<- struct{}) {
	url := fmt.Sprintf("%s%s%s", c.baseURL, symbol, depth)
	return c.subscribeAggTrade(url, handler)
}
