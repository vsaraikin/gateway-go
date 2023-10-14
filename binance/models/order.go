package models

type Side string

type TestOrder struct{}

// OrderRequest represents an order to be sent to Binance API.
type OrderRequest struct {
	Symbol                  string  `json:"symbol"`
	Side                    string  `json:"side"`
	Type                    string  `json:"type"`
	TimeInForce             string  `json:"timeInForce,omitempty"`
	Quantity                string  `json:"quantity,omitempty"`
	QuoteOrderQty           string  `json:"quoteOrderQty,omitempty"`
	Price                   string  `json:"price,omitempty"`
	NewClientOrderID        string  `json:"newClientOrderId,omitempty"`
	StopPrice               string  `json:"stopPrice,omitempty"`
	IcebergQty              string  `json:"icebergQty,omitempty"`
	NewOrderRespType        string  `json:"newOrderRespType,omitempty"`
	RecvWindow              string  `json:"recvWindow,omitempty"`
	Timestamp               string  `json:"timestamp"`
	StrategyID              int     `json:"strategyId,omitempty"`
	StrategyType            int     `json:"strategyType,omitempty"`
	TrailingDelta           int64   `json:"trailingDelta,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
}
