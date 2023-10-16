package models

type Side string

type OrderType string

type TestOrder struct{}

// OrderRequest represents an order to be sent to Binance API.
type OrderRequest struct {
	Symbol                  string  `url:"symbol"`
	Side                    string  `url:"side"`
	Type                    string  `url:"type"`
	TimeInForce             string  `url:"timeInForce,omitempty"`
	Quantity                int64   `url:"quantity,omitempty"`
	QuoteOrderQty           string  `url:"quoteOrderQty,omitempty"`
	Price                   string  `url:"price,omitempty"`
	NewClientOrderID        string  `url:"newClientOrderId,omitempty"`
	StopPrice               string  `url:"stopPrice,omitempty"`
	IcebergQty              string  `url:"icebergQty,omitempty"`
	NewOrderRespType        string  `url:"newOrderRespType,omitempty"`
	RecvWindow              int64   `url:"recvWindow,omitempty"`
	Timestamp               int64   `url:"timestamp"`
	StrategyID              int     `url:"strategyId,omitempty"`
	StrategyType            int     `url:"strategyType,omitempty"`
	TrailingDelta           int64   `url:"trailingDelta,omitempty"`
	SelfTradePreventionMode *string `url:"selfTradePreventionMode,omitempty"`
}

// OrderType and other enum values
const (
	BUY  = "BUY"
	SELL = "SELL"

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
