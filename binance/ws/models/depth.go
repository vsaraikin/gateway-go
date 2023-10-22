package models

import (
	"strconv"
)

type OrderBook struct {
	Price    float32
	Quantity float32
}

type DepthEventRaw struct {
	Event         string      `json:"e"`
	Time          int64       `json:"E"`
	Symbol        string      `json:"s"`
	LastUpdateID  int64       `json:"u"`
	FirstUpdateID int64       `json:"U"`
	Bids          [][2]string `json:"b"`
	Asks          [][2]string `json:"a"`
}

type DepthEvent struct {
	Event         string      `json:"e"`
	Time          int64       `json:"E"`
	Symbol        string      `json:"s"`
	LastUpdateID  int64       `json:"u"`
	FirstUpdateID int64       `json:"U"`
	Bids          []OrderBook `json:"b"`
	Asks          []OrderBook `json:"a"`
}

// Transform changes data structure where orderbook is `float`
func (event *DepthEventRaw) Transform() *DepthEvent {
	output := &DepthEvent{
		Event:         event.Event,
		Time:          event.Time,
		Symbol:        event.Symbol,
		LastUpdateID:  event.LastUpdateID,
		FirstUpdateID: event.FirstUpdateID,
	}

	for _, bid := range event.Bids {
		price, _ := strconv.ParseFloat(bid[0], 64)
		quantity, _ := strconv.ParseFloat(bid[1], 64)
		output.Bids = append(output.Bids, OrderBook{Price: float32(price), Quantity: float32(quantity)})
	}

	for _, ask := range event.Asks {
		price, _ := strconv.ParseFloat(ask[0], 64)
		quantity, _ := strconv.ParseFloat(ask[1], 64)
		output.Asks = append(output.Asks, OrderBook{Price: float32(price), Quantity: float32(quantity)})
	}

	return output
}
