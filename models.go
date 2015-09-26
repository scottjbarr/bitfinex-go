package bitfinex

// Ticker
//
// Sample response
//
//   {
//     'ask': '562.9999',
//     'timestamp': '1395552290.70933607',
//     'bid': '562.25',
//     'last_price': u'562.25',
//     'mid': u'562.62495'
//   }
type Ticker struct {
	LastTrade float32 `json:"last_price,string"`
	Bid       float32 `json:"bid,string"`
	Ask       float32 `json:"ask,string"`
}

// OrderBook represents the full order book returned by the API.
//
//   curl "https://api.bitfinex.com/v1/book/btcusd"
//
//   {
//     "bids":[
//       {"price":"561.1101","amount":"0.985","timestamp":"1395557729.0"},
//       {..}
//     ],
//     "asks":[
//       {"price":"562.9999","amount":"0.985","timestamp":"1395557711.0"},
//       {..}
//     ]
//  }
//
// Optional parameters
//
// - limit_bids (int): Optional. Limit the number of bids returned. May be 0 in
//   which case the array of bids is empty. Default is 50.
// - limit_asks (int): Optional. Limit the number of asks returned. May be 0 in
//   which case the array of asks is empty. Default is 50.
//
// eg.
//   curl "https://api.bitfinex.com/v1/book/btcusd?limit_bids=1&limit_asks=0"
type OrderBook struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

// Order represents the price and quanty of an individual Order, or the summary
// of multiple Orders (as in the case of an Order Book)
type Order struct {
	Price    float32 `json:"price,string"`
	Quantity float32 `json:"amount,string"`
}
