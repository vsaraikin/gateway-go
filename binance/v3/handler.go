package v3

import (
	"encoding/json"
	"fmt"
	"gateaway/binance/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

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

type BinanceClient struct {
	APIKey  string
	Secret  string
	BaseURL string
	client  http.Client
}

func NewBinanceClient(apiKey, secretKey string) *BinanceClient {
	return &BinanceClient{
		APIKey:  apiKey,
		Secret:  secretKey,
		BaseURL: "https://api.binance.com",
		client:  http.Client{},
	}
}

func (c *BinanceClient) executeRequest(method, endpoint string, body io.Reader, target interface{}, sign bool, params interface{}) error {
	m := map[string]interface{}{
		"symbol":     "BTCUSDT",
		"side":       "BUY",
		"quantity":   1,
		"type":       "MARKET",
		"recvWindow": 10000,
		"timestamp":  time.Now().UnixMilli(),
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	// Converting map to url.Values
	values := url.Values{}
	for key, value := range m {
		values.Add(key, fmt.Sprintf("%v", value))
	}

	// Encoding map as URL parameters
	u.RawQuery = values.Encode()

	if sign {
		u.RawQuery = fmt.Sprintf("%s&signature=%s", u.RawQuery, signature(u.RawQuery, c.Secret))
	}

	req, err := http.NewRequest(method, u.String(), body) // passing 'body' instead of 'nil'
	if err != nil {
		return err
	}

	if sign {
		req.Header.Add("X-MBX-APIKEY", c.APIKey)
	}

	fmt.Println(u.RequestURI())

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

// ––––––––––– MARKET DATA –––––––––––

func (c *BinanceClient) getExchangeInfo(url string) (*models.ExchangeInfo, error) {
	info := &models.ExchangeInfo{}
	params := make(map[string]string)
	err := c.executeRequest("GET", url, nil, info, false, params)

	if err != nil {
		return nil, err
	}

	return info, nil
}

// GetExchangeInfo Current exchange trading rules and symbol information
func (c *BinanceClient) GetExchangeInfo() (*models.ExchangeInfo, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, exchangeInfo)
	return c.getExchangeInfo(url)
}

// ––––––––––– SPOT TRADING –––––––––––

func (c *BinanceClient) testNewOrder(url string, r interface{}) error {
	//params, err := parseModel(r)
	//if err != nil {
	//	return err
	//}

	err := c.executeRequest("POST", url, nil, struct{}{}, true, r)

	if err != nil {
		return err
	}

	return nil
}

// NewOrderTest
// Test new order creation and signature/recvWindow long.
// Creates and validates a new order but does not send it into the matching engine.
func (c *BinanceClient) NewOrderTest(request models.OrderRequest) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, testNewOrder)
	return c.testNewOrder(url, request)
}
