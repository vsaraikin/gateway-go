package models

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type DepthResponse struct {
	LastUpdateId int     `json:"lastUpdateId"`
	Bids         []Order `json:"bids"`
	Asks         []Order `json:"asks"`
}

type Order struct {
	Price    decimal.Decimal
	Quantity decimal.Decimal
}

func (o *Order) UnmarshalJSON(data []byte) error {
	var order [2]decimal.Decimal

	if err := json.Unmarshal(data, &order); err != nil {
		return err
	}

	o.Price = order[0]
	o.Quantity = order[1]

	return nil
}

type DepthRequest struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit,omitempty"`
}

func (r DepthRequest) Validate() error {
	if r.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	return nil
}
