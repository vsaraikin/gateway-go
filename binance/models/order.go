package models

import (
	"errors"
	"strings"
)

type Side string

type OrderType string

type TestOrder struct{}

// OrderRequest represents an order to be sent to Binance API.
type OrderRequest struct {
	Symbol                  string  `url:"symbol"`
	Side                    string  `url:"side"`
	Type                    string  `url:"type"`
	TimeInForce             string  `url:"timeInForce,omitempty"`
	Quantity                float32 `url:"quantity"`
	QuoteOrderQty           float32 `url:"quoteOrderQty,omitempty"`
	Price                   float32 `url:"price,omitempty"`
	NewClientOrderID        string  `url:"newClientOrderId,omitempty"`
	StopPrice               float32 `url:"stopPrice,omitempty"`
	IcebergQty              float32 `url:"icebergQty,omitempty"`
	NewOrderRespType        string  `url:"newOrderRespType,omitempty"`
	RecvWindow              int64   `url:"recvWindow,omitempty"`
	Timestamp               int64   `url:"timestamp"`
	StrategyID              int     `url:"strategyId,omitempty"`
	StrategyType            int     `url:"strategyType,omitempty"`
	TrailingDelta           int64   `url:"trailingDelta,omitempty"`
	SelfTradePreventionMode *string `url:"selfTradePreventionMode,omitempty"`
}

// Validate request TODO: Add better validation
func (o *OrderRequest) Validate() error {
	// Required fields
	if o.Symbol == "" {
		return errors.New("Symbol is required")
	}
	if o.Side == "" {
		return errors.New("Side is required")
	}
	if o.Type == "" {
		return errors.New("Type is required")
	}
	if o.Timestamp == 0 {
		return errors.New("Timestamp is required")
	}

	// Validate Side values
	validSides := []string{"BUY", "SELL"}
	if !stringInSlice(o.Side, validSides) {
		return errors.New("Invalid Side value")
	}

	// Validate Type values
	validTypes := []string{"LIMIT", "MARKET", "STOP_LOSS", "STOP_LOSS_LIMIT", "TAKE_PROFIT", "TAKE_PROFIT_LIMIT", "LIMIT_MAKER"}
	if !stringInSlice(o.Type, validTypes) {
		return errors.New("Invalid Type value")
	}

	// Validate TimeInForce if present
	if o.TimeInForce != "" {
		validTimeInForce := []string{"GTC", "IOC", "FOK"}
		if !stringInSlice(o.TimeInForce, validTimeInForce) {
			return errors.New("Invalid TimeInForce value")
		}
	}

	// Validate Quantity if present (should be greater than 0)
	if o.Quantity != 0 && o.Quantity <= 0 {
		return errors.New("quantity should be greater than 0")
	}

	return nil
}

// Helper function to check if a string exists in a slice
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

// TODO: Add better types for `enum` values e.g. order – BUY/SELL, LIMIT/MARKET, etc.

type OrderResponseAck struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	OrderListId   int    `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

type OrderResponseResult struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type OrderResponseFull struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	Fills                   []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int    `json:"tradeId"`
	} `json:"fills"`
}
