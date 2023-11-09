package models

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
)

type KlineRequest struct {
	Symbol    string `url:"symbol"`
	Interval  string `url:"interval"`
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

func (kr *KlineRequest) Validate() error {
	if kr.Symbol == "" {
		return errors.New("symbol is required")
	}
	if kr.Interval == "" {
		return errors.New("interval is required")
	}

	return nil
}

type KlineResponse struct {
	OpenTime         int64
	OpenPrice        decimal.Decimal
	HighPrice        decimal.Decimal
	LowPrice         decimal.Decimal
	ClosePrice       decimal.Decimal
	Volume           decimal.Decimal
	CloseTime        int64
	QuoteVolume      decimal.Decimal
	NumberOfTrades   int
	TakerAssetVolume decimal.Decimal
	TakerQuoteVolume decimal.Decimal
	Ignore           string
}

func (k *KlineResponse) UnmarshalJSON(b []byte) error {
	var tmp interface{}
	var err error

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	k.OpenTime = int64(tmp.([]interface{})[0].(float64))
	k.OpenPrice, err = decimal.NewFromString(tmp.([]interface{})[1].(string))
	k.HighPrice, err = decimal.NewFromString(tmp.([]interface{})[2].(string))
	k.LowPrice, err = decimal.NewFromString(tmp.([]interface{})[3].(string))
	k.ClosePrice, err = decimal.NewFromString(tmp.([]interface{})[4].(string))
	k.Volume, err = decimal.NewFromString(tmp.([]interface{})[5].(string))
	closeInNano := tmp.([]interface{})[6].(float64)
	k.CloseTime = int64(closeInNano)
	k.QuoteVolume, err = decimal.NewFromString(tmp.([]interface{})[7].(string))
	k.NumberOfTrades = int(tmp.([]interface{})[8].(float64))
	k.TakerAssetVolume, err = decimal.NewFromString(tmp.([]interface{})[9].(string))
	k.TakerQuoteVolume, err = decimal.NewFromString(tmp.([]interface{})[10].(string))

	return err
}
