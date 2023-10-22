package models

type ExchangeInfo struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
	} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                          string        `json:"symbol"`
		Status                          string        `json:"status"`
		BaseAsset                       string        `json:"baseAsset"`
		BaseAssetPrecision              int           `json:"baseAssetPrecision"`
		QuoteAsset                      string        `json:"quoteAsset"`
		QuotePrecision                  int           `json:"quotePrecision"`
		QuoteAssetPrecision             int           `json:"quoteAssetPrecision"`
		OrderTypes                      []string      `json:"orderTypes"`
		IcebergAllowed                  bool          `json:"icebergAllowed"`
		OcoAllowed                      bool          `json:"ocoAllowed"`
		QuoteOrderQtyMarketAllowed      bool          `json:"quoteOrderQtyMarketAllowed"`
		AllowTrailingStop               bool          `json:"allowTrailingStop"`
		CancelReplaceAllowed            bool          `json:"cancelReplaceAllowed"`
		IsSpotTradingAllowed            bool          `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed          bool          `json:"isMarginTradingAllowed"`
		Filters                         []interface{} `json:"filters"`
		Permissions                     []string      `json:"permissions"`
		DefaultSelfTradePreventionMode  string        `json:"defaultSelfTradePreventionMode"`
		AllowedSelfTradePreventionModes []string      `json:"allowedSelfTradePreventionModes"`
	} `json:"symbols"`
}
