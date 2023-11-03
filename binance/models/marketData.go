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
