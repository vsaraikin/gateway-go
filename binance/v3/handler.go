package v3

import (
	"encoding/json"
	"fmt"
	"gateaway/binance/models"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	urlib "net/url"
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
	// Parse the base URL
	u, err := urlib.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	q, _ := query.Values(params) // use google lib to crawl map into url params

	u.RawQuery = q.Encode()

	if sign {
		u.RawQuery = fmt.Sprintf("%s&signature=%s", u.RawQuery, signature(u.RawQuery, c.Secret))
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	fmt.Println(u.String())

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

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

// ––––––––––– MARKET DATA –––––––––––

func (c *BinanceClient) getExchangeInfo(url string) (*models.ExchangeInfo, error) {
	response := &models.ExchangeInfo{}
	var params interface{} // no params needed
	err := c.executeRequest(http.MethodGet, url, nil, response, false, params)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetExchangeInfo Current exchange trading rules and symbol information
func (c *BinanceClient) GetExchangeInfo() (*models.ExchangeInfo, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, exchangeInfo)
	return c.getExchangeInfo(url)
}

// ––––––––––– SPOT TRADING –––––––––––

func (c *BinanceClient) testNewOrder(url string, r models.OrderRequest) error {
	err := r.Validate()
	if err != nil {
		return err
	}

	err = c.executeRequest(http.MethodPost, url, nil, struct{}{}, true, r)
	if err != nil {
		return err
	}

	return nil
}

// NewOrderTest
// Test new order creation and signature/recvWindow long.
// Creates and validates a new order but does not send it into the matching engine.
func (c *BinanceClient) NewOrderTest(r models.OrderRequest) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, testNewOrder)
	return c.testNewOrder(url, r)
}

func (c *BinanceClient) newOrder(url string, r models.OrderRequest) (models.OrderResponseAck, error) {
	return models.OrderResponseAck{}, nil
}

func (c *BinanceClient) NewOrder(r models.OrderRequest) (models.OrderResponseAck, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, testNewOrder)
	return c.newOrder(url, r)
}
