package bitfinex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	prefix         = "/v1"
	timeout        = 5 * time.Second
	tickerPath     = "ticker/%v"
	orderBookPath  = "book/%v"
	BTCUSD         = "btcusd"
	orderBookLimit = 5
)

// Client communicates with the API
type Client struct {
	Host       string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		Host:       "https://api.bitfinex.com",
		HTTPClient: &http.Client{Timeout: timeout},
	}
}

// url returns a full URL for a resource
func (c *Client) url(resource string) string {
	return fmt.Sprintf("%v%v/%v", c.Host, prefix, resource)
}

// get a response from a URL
func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.HTTPClient.Get(url)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetTicker returns a Ticker from the API
func (c *Client) GetTicker(symbol string) (*Ticker, error) {
	var body []byte
	var err error

	if body, err = c.get(c.url(fmt.Sprintf(tickerPath, symbol))); err != nil {
		return nil, err
	}

	var ticker Ticker

	if err = json.Unmarshal(body, &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}

// GetOrderBook returns the OrderBook from the API
func (c *Client) GetOrderBook(symbol string) (*OrderBook, error) {
	var body []byte
	var err error

	opts := fmt.Sprintf("?limit_bids=%v&limit_asks=%v",
		orderBookLimit,
		orderBookLimit)
	path := fmt.Sprintf(orderBookPath, symbol) + opts

	if body, err = c.get(c.url(path)); err != nil {
		return nil, err
	}

	orderBook := OrderBook{}

	if err := json.Unmarshal(body, &orderBook); err != nil {
		return nil, err
	}

	return &orderBook, nil
}
