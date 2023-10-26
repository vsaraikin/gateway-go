package v3

import (
	"encoding/json"
	"fmt"
	"gateaway/binance/models"
	//"github.com/charmbracelet/log"
	"github.com/rs/zerolog/log"

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
		return err
	}

	q, err := query.Values(params)
	if err != nil {
		return err
	}

	u.RawQuery = q.Encode()

	if sign {
		u.RawQuery = fmt.Sprintf("%s&signature=%s", u.RawQuery, signature(u.RawQuery, c.Secret))
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	log.Info().Msg(fmt.Sprintf("Requested %s %s", method, u.String()))

	if sign {
		req.Header.Add("X-MBX-APIKEY", c.APIKey)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d\n%s", response.StatusCode, data)
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
	url := c.buildURL(exchangeInfo)
	return c.getExchangeInfo(url)
}

func (c *BinanceClient) getDepth(url string, params models.DepthRequest) (*models.DepthResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	response := &models.DepthResponse{}
	err = c.executeRequest(http.MethodGet, url, nil, response, false, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *BinanceClient) GetDepth(r models.DepthRequest) (*models.DepthResponse, error) {
	url := c.buildURL(depth)
	return c.getDepth(url, r)
}

// ––––––––––– SPOT TRADING –––––––––––

// NewOrderTest
// Test new order creation and signature/recvWindow long.
// Creates and validates a new order but does not send it into the matching engine.
func (c *BinanceClient) NewOrderTest(r models.OrderRequest) (*models.OrderResponseFull, error) {
	url := c.buildURL(testNewOrder)
	return c.newOrder(url, r)
}

func (c *BinanceClient) NewOrder(r models.OrderRequest) (*models.OrderResponseFull, error) {
	url := c.buildURL(newOrder)
	return c.newOrder(url, r)
}

func (c *BinanceClient) newOrder(url string, params models.OrderRequest) (*models.OrderResponseFull, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	// When making the API call, you can specify which response type you want by setting
	// the newOrderRespType parameter to either ACK, RESULT, or FULL.
	// If you don't specify a type, the default is RESULT.
	response := &models.OrderResponseFull{}
	err = c.executeRequest(http.MethodPost, url, nil, response, true, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *BinanceClient) CancelOrder(r models.OrderCancelRequest) (*models.OrderCancelResponse, error) {
	url := c.buildURL(newOrder)
	return c.cancelOrder(url, r)
}

func (c *BinanceClient) cancelOrder(url string, params models.OrderCancelRequest) (*models.OrderCancelResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	response := &models.OrderCancelResponse{}
	err = c.executeRequest(http.MethodDelete, url, nil, response, true, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *BinanceClient) cancelAllOpenOrders(url string, params models.CancelAllOrdersRequest) (*models.CancelAllOrdersResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	response := &models.CancelAllOrdersResponse{}
	err = c.executeRequest(http.MethodDelete, url, nil, response, true, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *BinanceClient) CancelAllOpenOrders(r models.CancelAllOrdersRequest) (*models.CancelAllOrdersResponse, error) {
	url := c.buildURL(openOrders)
	return c.cancelAllOpenOrders(url, r)
}

//func (c *BinanceClient) queryOrder(url string, params)
