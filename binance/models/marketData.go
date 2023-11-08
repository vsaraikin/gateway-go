package models

import "errors"

type TradesRequest struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit,omitempty"`
}

func (r *TradesRequest) Validate() error {
	if r.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

type TradesResponse struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type AggTradeRequest struct {
	Symbol    string `url:"symbol"`
	FromID    int64  `url:"fromId,omitempty"`
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

func (t *AggTradeRequest) Validate() error {
	if t.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

type AggTradesResponse struct {
	A  int    `json:"a"`
	P  string `json:"p"`
	Q  string `json:"q"`
	F  int    `json:"f"`
	L  int    `json:"l"`
	T  int64  `json:"T"`
	M  bool   `json:"m"`
	M1 bool   `json:"M"`
}
