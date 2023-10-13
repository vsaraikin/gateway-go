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
	baseURL string
	client  http.Client
}

func NewBinanceClient(apiKey, secretKey string) *BinanceClient {
	return &BinanceClient{
		APIKey:  apiKey,
		Secret:  secretKey,
		baseURL: "https://api.binance.com",
		client:  http.Client{},
	}
}

func (c *BinanceClient) executeRequest(method, url string, body io.Reader, target interface{}, sign bool) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if sign {

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

func (c *BinanceClient) testNewOrder(url string) error {
	err := c.executeRequest("POST", url, nil, struct{}{}, true)

	if err != nil {
		return err
	}

	return nil
}

func (c *BinanceClient) getExchangeInfo(url string) (*models.ExchangeInfo, error) {
	info := &models.ExchangeInfo{}
	err := c.executeRequest("GET", url, nil, info, false)

	if err != nil {
		return nil, err
	}

	return info, nil
}

// NewOrderTest
// Test new order creation and signature/recvWindow long.
// Creates and validates a new order but does not send it into the matching engine.
func (c *BinanceClient) NewOrderTest() error {
	url := fmt.Sprintf("%s%s", c.baseURL, testNewOrder)
	return c.testNewOrder(url)
}

// GetExchangeInfo Current exchange trading rules and symbol information
func (c *BinanceClient) GetExchangeInfo() (*models.ExchangeInfo, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, exchangeInfo)
	return c.getExchangeInfo(url)
}
