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

type CancelReplaceRequest struct {
	Symbol                  string  `url:"symbol"`
	Side                    string  `url:"side"`
	Type                    string  `url:"type"`
	CancelReplaceMode       string  `url:"cancelReplaceMode"`
	TimeInForce             string  `url:"timeInForce,omitempty"`
	Quantity                float64 `url:"quantity,omitempty"`
	QuoteOrderQty           float64 `url:"quoteOrderQty,omitempty"`
	Price                   float64 `url:"price,omitempty"`
	CancelNewClientOrderId  string  `url:"cancelNewClientOrderId,omitempty"`
	CancelOrigClientOrderId string  `url:"cancelOrigClientOrderId,omitempty"`
	CancelOrderId           int64   `url:"cancelOrderId,omitempty"`
	NewClientOrderId        string  `url:"newClientOrderId,omitempty"`
	StrategyId              int     `url:"strategyId,omitempty"`
	StrategyType            int     `url:"strategyType,omitempty"`
	StopPrice               float64 `url:"stopPrice,omitempty"`
	TrailingDelta           int64   `url:"trailingDelta,omitempty"`
	IcebergQty              float64 `url:"icebergQty,omitempty"`
	NewOrderRespType        string  `url:"newOrderRespType,omitempty"`
	SelfTradePreventionMode string  `url:"selfTradePreventionMode,omitempty"`
	CancelRestrictions      string  `url:"cancelRestrictions,omitempty"`
	RecvWindow              int64   `url:"recvWindow,omitempty"`
	Timestamp               int64   `url:"timestamp"`
}

func (req *CancelReplaceRequest) Validate() error {
	// Validate mandatory fields
	if req.Symbol == "" {
		return errors.New("symbol is mandatory")
	}
	if req.Side != "BUY" && req.Side != "SELL" {
		return errors.New("side must be either BUY or SELL")
	}
	if req.Type == "" {
		return errors.New("type is mandatory")
	}
	if req.CancelReplaceMode != "STOP_ON_FAILURE" && req.CancelReplaceMode != "ALLOW_FAILURE" {
		return errors.New("cancelReplaceMode must be either STOP_ON_FAILURE or ALLOW_FAILURE")
	}
	if req.Timestamp == 0 {
		return errors.New("timestamp is mandatory")
	}

	if req.StrategyType != 0 && req.StrategyType < 1000000 {
		return errors.New("strategyType must be greater than or equal to 1000000")
	}

	if req.RecvWindow != 0 && req.RecvWindow > 60000 {
		return errors.New("recvWindow must be less than or equal to 60000")
	}

	if req.CancelOrigClientOrderId == "" && req.CancelOrderId == 0 {
		return errors.New("either cancelOrigClientOrderId or cancelOrderId must be provided")
	}

	if req.NewOrderRespType != "" {
		switch req.NewOrderRespType {
		case "ACK", "RESULT", "FULL":
			// valid
		default:
			return errors.New("newOrderRespType must be either ACK, RESULT, or FULL")
		}
	}

	if req.SelfTradePreventionMode != "" {
		switch req.SelfTradePreventionMode {
		case "EXPIRE_TAKER", "EXPIRE_MAKER", "EXPIRE_BOTH", "NONE":
			// valid
		default:
			return errors.New("selfTradePreventionMode is not valid")
		}
	}

	if req.CancelRestrictions != "" {
		switch req.CancelRestrictions {
		case "ONLY_NEW", "ONLY_PARTIALLY_FILLED":
			// valid
		default:
			return errors.New("cancelRestrictions is not valid")
		}
	}

	return nil
}

type CancelReplaceResponse struct {
	Code           int64  `json:"code,omitempty"`
	Msg            string `json:"msg,omitempty"`
	CancelResult   string `json:"cancelResult,omitempty"`
	NewOrderResult string `json:"newOrderResult,omitempty"`
	CancelResponse struct {
		Code                    int    `json:"code,omitempty"`
		Msg                     string `json:"msg,omitempty"`
		Symbol                  string `json:"symbol,omitempty"`
		OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
		OrderId                 int64  `json:"orderId,omitempty"`
		OrderListId             int64  `json:"orderListId,omitempty"`
		ClientOrderId           string `json:"clientOrderId,omitempty"`
		Price                   string `json:"price,omitempty"`
		OrigQty                 string `json:"origQty,omitempty"`
		ExecutedQty             string `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
		Status                  string `json:"status,omitempty"`
		TimeInForce             string `json:"timeInForce,omitempty"`
		Type                    string `json:"type,omitempty"`
		Side                    string `json:"side,omitempty"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	} `json:"cancelResponse,omitempty"`
	NewOrderResponse struct {
		Code                    int64    `json:"code,omitempty"`
		Msg                     string   `json:"msg,omitempty"`
		Symbol                  string   `json:"symbol,omitempty"`
		OrderId                 int64    `json:"orderId,omitempty"`
		OrderListId             int64    `json:"orderListId,omitempty"`
		ClientOrderId           string   `json:"clientOrderId,omitempty"`
		TransactTime            uint64   `json:"transactTime,omitempty"`
		Price                   string   `json:"price,omitempty"`
		OrigQty                 string   `json:"origQty,omitempty"`
		ExecutedQty             string   `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
		Status                  string   `json:"status,omitempty"`
		TimeInForce             string   `json:"timeInForce,omitempty"`
		Type                    string   `json:"type,omitempty"`
		Side                    string   `json:"side,omitempty"`
		Fills                   []string `json:"fills,omitempty"`
		SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
	} `json:"newOrderResponse,omitempty"`
	Data struct {
		CancelResult   string `json:"cancelResult,omitempty"`
		NewOrderResult string `json:"newOrderResult,omitempty"`
		CancelResponse struct {
			Code                    int64  `json:"code,omitempty"`
			Msg                     string `json:"msg,omitempty"`
			Symbol                  string `json:"symbol,omitempty"`
			OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
			OrderId                 int64  `json:"orderId,omitempty"`
			OrderListId             int64  `json:"orderListId,omitempty"`
			ClientOrderId           string `json:"clientOrderId,omitempty"`
			Price                   string `json:"price,omitempty"`
			OrigQty                 string `json:"origQty,omitempty"`
			ExecutedQty             string `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
			Status                  string `json:"status,omitempty"`
			TimeInForce             string `json:"timeInForce,omitempty"`
			Type                    string `json:"type,omitempty"`
			Side                    string `json:"side,omitempty"`
			SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
		} `json:"cancelResponse,omitempty"`
		NewOrderResponse struct {
			Code                    int64    `json:"code,omitempty"`
			Msg                     string   `json:"msg,omitempty"`
			Symbol                  string   `json:"symbol,omitempty"`
			OrderId                 int64    `json:"orderId,omitempty"`
			OrderListId             int64    `json:"orderListId,omitempty"`
			ClientOrderId           string   `json:"clientOrderId,omitempty"`
			TransactTime            uint64   `json:"transactTime,omitempty"`
			Price                   string   `json:"price,omitempty"`
			OrigQty                 string   `json:"origQty,omitempty"`
			ExecutedQty             string   `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
			Status                  string   `json:"status,omitempty"`
			TimeInForce             string   `json:"timeInForce,omitempty"`
			Type                    string   `json:"type,omitempty"`
			Side                    string   `json:"side,omitempty"`
			Fills                   []string `json:"fills,omitempty"`
			SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
		} `json:"newOrderResponse"`
	} `json:"data,omitempty"`
}

type OpenOrdersRequest struct {
	Symbol     string `url:"symbol,omitempty"`
	RecvWindow *int64 `url:"recvWindow,omitempty"`
	Timestamp  int64  `url:"timestamp"`
}

func (o *OpenOrdersRequest) Validate() error {
	if o.Timestamp <= 0 {
		return errors.New("timestamp is required and must be a positive integer")
	}
	return nil
}

type OpenOrdersResponse struct {
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

type AllOpenOrdersRequest struct {
	Symbol     string `url:"symbol"`
	OrderID    *int64 `url:"orderId,omitempty"`
	StartTime  *int64 `url:"startTime,omitempty"`
	EndTime    *int64 `url:"endTime,omitempty"`
	Limit      *int   `url:"limit,omitempty"`
	RecvWindow *int64 `url:"recvWindow,omitempty"`
	Timestamp  int64  `url:"timestamp"`
}

func (o *AllOpenOrdersRequest) Validate() error {
	if o.Symbol == "" {
		return errors.New("symbol is required and cannot be empty")
	}

	return nil
}

type AllOpenOrdersResponse struct {
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
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type NewOCORequest struct {
	Symbol                  string   `url:"symbol"`
	ListClientOrderId       *string  `url:"listClientOrderId,omitempty"`
	Side                    Side     `url:"side"`
	Quantity                float64  `url:"quantity"`
	LimitClientOrderId      *string  `url:"limitClientOrderId,omitempty"`
	LimitStrategyId         *int     `url:"limitStrategyId,omitempty"`
	LimitStrategyType       *int     `url:"limitStrategyType,omitempty"`
	Price                   float64  `url:"price"`
	LimitIcebergQty         *float64 `url:"limitIcebergQty,omitempty"`
	TrailingDelta           *int64   `url:"trailingDelta,omitempty"`
	StopClientOrderId       *string  `url:"stopClientOrderId,omitempty"`
	StopPrice               float64  `url:"stopPrice"`
	StopStrategyId          *int     `url:"stopStrategyId,omitempty"`
	StopStrategyType        *int     `url:"stopStrategyType,omitempty"`
	StopLimitPrice          *float64 `url:"stopLimitPrice,omitempty"`
	StopIcebergQty          *float64 `url:"stopIcebergQty,omitempty"`
	StopLimitTimeInForce    string   `url:"stopLimitTimeInForce,omitempty"`
	NewOrderRespType        string   `url:"newOrderRespType,omitempty"`
	SelfTradePreventionMode string   `url:"selfTradePreventionMode,omitempty"`
	RecvWindow              *int64   `url:"recvWindow,omitempty"`
	Timestamp               int64    `url:"timestamp"`
}

func (r *NewOCORequest) Validate() error {
	if r.Symbol == "" {
		return errors.New("symbol is required and cannot be empty")
	}

	if r.Side != "BUY" && r.Side != "SELL" {
		return errors.New("side is required and must be either BUY or SELL")
	}

	if r.Quantity <= 0 {
		return errors.New("quantity is required and must be greater than 0")
	}

	if r.Price <= 0 {
		return errors.New("price is required and must be greater than 0")
	}

	if r.StopPrice <= 0 {
		return errors.New("stopPrice is required and must be greater than 0")
	}

	if r.LimitStrategyType != nil && *r.LimitStrategyType < 1000000 {
		return errors.New("limitStrategyType cannot be less than 1000000")
	}

	if r.StopStrategyType != nil && *r.StopStrategyType < 1000000 {
		return errors.New("stopStrategyType cannot be less than 1000000")
	}

	if r.RecvWindow != nil && *r.RecvWindow > 60000 {
		return errors.New("recvWindow cannot be greater than 60000")
	}

	if r.Timestamp <= 0 {
		return errors.New("timestamp is required and must be a positive integer")
	}

	// Price and quantity restrictions based on side
	if r.Side == "SELL" && (r.Price <= r.StopPrice) {
		return errors.New("for a SELL order, limit price must be greater than the stop price")
	}

	if r.Side == "BUY" && (r.Price >= r.StopPrice) {
		return errors.New("for a BUY order, limit price must be less than the stop price")
	}

	// If stopLimitPrice is provided, stopLimitTimeInForce must also be provided
	if r.StopLimitPrice != nil && r.StopLimitTimeInForce == "" {
		return errors.New("if stopLimitPrice is provided, stopLimitTimeInForce is required")
	}

	return nil
}

type NewOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
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
		StopPrice               string `json:"stopPrice,omitempty"`
		WorkingTime             int64  `json:"workingTime"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"orderReports"`
}

type CancelOCORequest struct {
	Symbol            string  `url:"symbol"`
	OrderListID       *int    `url:"orderListId,omitempty"`
	ListClientOrderID *string `url:"listClientOrderId,omitempty"`
	NewClientOrderID  *string `url:"newClientOrderId,omitempty"`
	RecvWindow        *int64  `url:"recvWindow,omitempty"`
	Timestamp         int64   `url:"timestamp"`
}

func (r *CancelOCORequest) Validate() error {
	if r.Symbol == "" {
		return errors.New("symbol is required")
	}

	if r.OrderListID == nil && r.ListClientOrderID == nil {
		return errors.New("either orderListId or listClientOrderId must be provided")
	}

	if r.RecvWindow != nil && *r.RecvWindow > 60000 {
		return errors.New("recvWindow cannot be greater than 60000")
	}

	if r.Timestamp <= 0 {
		return errors.New("timestamp must be a positive integer and is required")
	}

	return nil
}

type CancelOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
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
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	} `json:"orderReports"`
}

type GetOCORequest struct {
	OrderListID       *int    `url:"orderListId,omitempty"`
	OrigClientOrderID *string `url:"origClientOrderId,omitempty"`
	RecvWindow        *int    `url:"recvWindow,omitempty"`
	Timestamp         int64   `url:"timestamp"`
}

func (p *GetOCORequest) Validate() error {
	if (p.OrderListID == nil) == (p.OrigClientOrderID == nil) {
		return errors.New("either orderListId or origClientOrderId must be provided, but not both")
	}

	if p.RecvWindow != nil && *p.RecvWindow > 60000 {
		return errors.New("recvWindow cannot be greater than 60000")
	}

	if p.Timestamp <= 0 {
		return errors.New("timestamp must be a positive integer and is required")
	}

	return nil
}

type GetOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

type AllOCOListRequest struct {
	FromID     *int64 `url:"fromId,omitempty"`
	StartTime  *int64 `url:"startTime,omitempty"`
	EndTime    *int64 `url:"endTime,omitempty"`
	Limit      *int   `url:"limit,omitempty"`
	RecvWindow *int64 `url:"recvWindow,omitempty"`
	Timestamp  int64  `url:"timestamp"`
}

func (r *AllOCOListRequest) Validate() error {
	if r.FromID != nil && (r.StartTime != nil || r.EndTime != nil) {
		return errors.New("if fromId is supplied, neither startTime nor endTime can be provided")
	}

	return nil
}

type AllOCOListResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}
