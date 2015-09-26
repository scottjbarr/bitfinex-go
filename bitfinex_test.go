package bitfinex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

// Mock out HTTP requests.
//
// Pinched from http://keighl.com/post/mocking-http-responses-in-golang/
// Thanks, Kyle Truscott (@keighl)!
func httpMock(code int,
	body string) (*httptest.Server, *Client) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport, Timeout: 5 * time.Second}

	// setting the host in the Client so I don't need to totally fake out
	// the TLS config
	client := &Client{
		Host:       "http://api.bitfinex.com",
		HTTPClient: httpClient,
	}

	return server, client
}

// Test helper. Thanks again, @keighl
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a))
	}
}

func TestGetTicker(t *testing.T) {

	//   {
	//     'ask': '562.9999',
	//     'timestamp': '1395552290.70933607',
	//     'bid': '562.25',
	//     'last_price': u'562.25',
	//     'mid': u'562.62495'
	//   }

	body := `{"ask":"562.9999","timestamp":"1395552290.70933607","bid":"562.25","last_price":"562.26","mid":"562.62495"}`
	server, client := httpMock(200, body)
	defer server.Close()

	ticker, err := client.GetTicker(BTCUSD)

	if err != nil {
		t.Errorf("GetTicker : %v", err)
	}

	expect(t, ticker.LastTrade, float32(562.26))
	expect(t, ticker.Bid, float32(562.25))
	expect(t, ticker.Ask, float32(562.9999))
}

func TestGetOrderBook(t *testing.T) {
	body := `{"bids":[{"price":"234.81","amount":"7.0957","timestamp":"1443269406.0"}],"asks":[{"price":"234.96","amount":"0.37","timestamp":"1443269216.0"}]}`
	server, client := httpMock(200, body)
	defer server.Close()

	orderBook, err := client.GetOrderBook(BTCUSD)

	if err != nil {
		t.Errorf("GetDepth : %v", err)
	}

	expect(t, len(orderBook.Asks), 1)
	expect(t, orderBook.Asks[0].Price, float32(234.96))
	expect(t, orderBook.Asks[0].Quantity, float32(0.37))
	expect(t, len(orderBook.Bids), 1)
	expect(t, orderBook.Bids[0].Price, float32(234.81))
	expect(t, orderBook.Bids[0].Quantity, float32(7.0957))

}
