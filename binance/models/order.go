package models

type TestOrder struct{}

// OrderType and other enum values
const (
	// Order types
	LIMIT             = "LIMIT"
	MARKET            = "MARKET"
	STOP_LOSS         = "STOP_LOSS"
	STOP_LOSS_LIMIT   = "STOP_LOSS_LIMIT"
	TAKE_PROFIT       = "TAKE_PROFIT"
	TAKE_PROFIT_LIMIT = "TAKE_PROFIT_LIMIT"
	LIMIT_MAKER       = "LIMIT_MAKER"

	// Time in force
	GTC = "GTC"
	IOC = "IOC"
	FOK = "FOK"

	// New Order Response Type
	ACK    = "ACK"
	RESULT = "RESULT"
	FULL   = "FULL"

	// Self Trade Prevention Mode
	EXPIRE_TAKER = "EXPIRE_TAKER"
	EXPIRE_MAKER = "EXPIRE_MAKER"
	EXPIRE_BOTH  = "EXPIRE_BOTH"
	NONE         = "NONE"
)

// OrderRequest represents an order to be sent to Binance API.
type OrderRequest struct {
	Symbol                  string  `json:"symbol"`
	Side                    string  `json:"side"`
	Type                    string  `json:"type"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Quantity                *string `json:"quantity,omitempty"`
	QuoteOrderQty           *string `json:"quoteOrderQty,omitempty"`
	Price                   *string `json:"price,omitempty"`
	NewClientOrderID        *string `json:"newClientOrderId,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	IcebergQty              *string `json:"icebergQty,omitempty"`
	NewOrderRespType        *string `json:"newOrderRespType,omitempty"`
	RecvWindow              *int64  `json:"recvWindow,omitempty"`
	Timestamp               int64   `json:"timestamp"`
	StrategyID              *int    `json:"strategyId,omitempty"`
	StrategyType            *int    `json:"strategyType,omitempty"`
	TrailingDelta           *int64  `json:"trailingDelta,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
}
