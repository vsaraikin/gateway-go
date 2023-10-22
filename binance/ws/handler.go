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
		go func() {
			<-done
			defer conn.Close()
		}()

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
			err = handler(message)
			if err != nil {
				fmt.Errorf(err.Error())
				// TODO: better handling
			}
		}
	}()

	return nil, done
}

type handlerEvent func(e *models.DepthEvent)

func (c *BinanceWsClient) serveDepth(url string, handler handlerEvent) (error, chan<- struct{}) {
	wsHandler := func(event []byte) error {
		depthEventRaw := new(models.DepthEventRaw)
		if err := json.Unmarshal(event, depthEventRaw); err != nil {
			fmt.Errorf("error json parsing %s", err.Error())
			// TODO: better handling
		}
		depthEvent := depthEventRaw.Transform()
		handler(depthEvent)
		return nil
	}
	return c.subscribe(url, wsHandler)
}

func (c *BinanceWsClient) SubscribeDepth(symbol string, handler handlerEvent) (error, chan<- struct{}) {
	url := fmt.Sprintf("%s%s%s", c.baseURL, symbol, depth)
	return c.serveDepth(url, handler)
}
