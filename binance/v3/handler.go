package v3

import (
	"encoding/json"
	"fmt"
	"gateaway/binance/models"
	"io"
	"io/ioutil"
	"net/http"
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

func (c *BinanceClient) executeRequest(method, url string, body io.Reader, target interface{}, sign bool, params map[string]interface{}) error {
	if sign {
		url = fmt.Sprintf("%s&signature=%s", url, signature(url, c.Secret))
	}
	// url.Values as params for request

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if sign {
		req.Header.Add("X-MBX-APIKEY", c.APIKey)
	}

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
	params := make(map[string]interface{})
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
	params, err := parseModel(r)
	if err != nil {
		return err
	}

	err = c.executeRequest("POST", url, nil, struct{}{}, true, params)

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
