package models

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

// NEW ORDER

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

// Validate request
func (o *OrderRequest) Validate() error {
	// Required fields
	if o.Symbol == "" {
		return errors.New("symbol is required")
	}
	if o.Side == "" {
		return errors.New("side is required")
	}
	if o.Type == "" {
		return errors.New("type is required")
	}
	if o.Timestamp == 0 {
		return errors.New("timestamp is required")
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
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
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

// CANCEL ORDER

type OrderCancelRequest struct {
	Symbol            string `url:"symbol" binding:"required"`
	OrderID           int64  `url:"orderId,omitempty"`
	OrigClientOrderID string `url:"origClientOrderId,omitempty"`
	NewClientOrderID  string `url:"newClientOrderId,omitempty"`
	CancelRestriction string `url:"cancelRestrictions,omitempty"`
	RecvWindow        int64  `url:"recvWindow,omitempty" binding:"omitempty,lt=60001"`
	Timestamp         int64  `url:"timestamp" binding:"required"`
}

func (o *OrderCancelRequest) Validate() error {
	if o.CancelRestriction != "ONLY_NEW" && o.CancelRestriction != "ONLY_PARTIALLY_FILLED" {
		return fmt.Errorf("incorrect cancelrestriction value")
	}
	return nil
}

type OrderCancelResponse struct {
	Symbol                  string          `json:"symbol"`
	OrigClientOrderId       string          `json:"origClientOrderId"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int64           `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	TransactTime            int64           `json:"transactTime"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 string          `json:"origQty"`
	ExecutedQty             string          `json:"executedQty"`
	CummulativeQuoteQty     string          `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

type CancelAllOrdersRequest struct {
	Symbol     string `url:"symbol" binding:"required"`
	RecvWindow int64  `url:"recvWindow,omitempty"`
	Timestamp  int64  `url:"timestamp" binding:"required"`
}

func (r CancelAllOrdersRequest) Validate() error {
	if r.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if r.Timestamp == 0 {
		return fmt.Errorf("timestamp cannot be empty")
	}
	return nil
}

type CancelAllOrdersResponse struct {
	Symbol                  string `json:"symbol"`
	OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
	OrderId                 int    `json:"orderId,omitempty"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId,omitempty"`
	TransactTime            int64  `json:"transactTime,omitempty"`
	Price                   string `json:"price,omitempty"`
	OrigQty                 string `json:"origQty,omitempty"`
	ExecutedQty             string `json:"executedQty,omitempty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty,omitempty"`
	Status                  string `json:"status,omitempty"`
	TimeInForce             string `json:"timeInForce,omitempty"`
	Type                    string `json:"type,omitempty"`
	Side                    string `json:"side,omitempty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	ContingencyType         string `json:"contingencyType,omitempty"`
	ListStatusType          string `json:"listStatusType,omitempty"`
	ListOrderStatus         string `json:"listOrderStatus,omitempty"`
	ListClientOrderId       string `json:"listClientOrderId,omitempty"`
	TransactionTime         int64  `json:"transactionTime,omitempty"`
	Orders                  []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders,omitempty"`
	OrderReports []struct {
		Symbol                  string `json:"symbol"`
		OrigClientOrderId       string `json:"origClientOrderId"`
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
		StopPrice               string `json:"stopPrice,omitempty"`
		IcebergQty              string `json:"icebergQty"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"orderReports,omitempty"`
}

type GetOrderRequest struct {
	Symbol            string `url:"symbol"`
	OrderID           int64  `url:"orderId,omitempty"`
	OrigClientOrderID string `url:"origClientOrderId,omitempty"`
	RecvWindow        int64  `url:"recvWindow,omitempty"`
	Timestamp         int64  `url:"timestamp"`
}

type GetOrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty"`
	Time                    int64  `json:"time"`
	UpdateTime              int64  `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             int64  `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

func (o *GetOrderRequest) Validate() error {
	if o.Symbol == "" {
		return errors.New("symbol is mandatory")
	}

	if o.Timestamp == 0 {
		return errors.New("timestamp is mandatory")
	}

	if o.OrderID == 0 {
		return fmt.Errorf("provide correct order id")
	}

	return nil
}
