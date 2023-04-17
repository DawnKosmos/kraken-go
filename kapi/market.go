package kapi

type GetServerTimeResult struct {
	Unixtime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}

type GetAssetInfoRequest struct {
	Asset string `url:"asset"`
}

type Asset struct {
	Aclass          string  `json:"aclass"`
	Altname         string  `json:"altname"`
	Decimals        float64 `json:"decimals"`
	DisplayDecimals int     `json:"display_decimals"`
	CollateralValue float64 `json:"collateral_value"`
	Status          string  `json:"status"`
}

type GetTradablePairRequest struct {
	Pair string `url:"pair"`
}

type TradableAsset struct {
	Altname            string       `json:"altname"`
	Wsname             string       `json:"wsname"`
	AclassBase         string       `json:"aclass_base"`
	Base               string       `json:"base"`
	AclassQuote        string       `json:"aclass_quote"`
	Quote              string       `json:"quote"`
	Lot                string       `json:"lot"`
	CostDecimals       int          `json:"cost_decimals"`
	PairDecimals       int          `json:"pair_decimals"`
	LotDecimals        int          `json:"lot_decimals"`
	LotMultiplier      int          `json:"lot_multiplier"`
	LeverageBuy        []int        `json:"leverage_buy"`
	LeverageSell       []int        `json:"leverage_sell"`
	Fees               [][2]float64 `json:"fees"`
	FeesMaker          [][2]float64 `json:"fees_maker"`
	FeeVolumeCurrency  string       `json:"fee_volume_currency"`
	MarginCall         int          `json:"margin_call"`
	MarginStop         int          `json:"margin_stop"`
	Ordermin           string       `json:"ordermin"`
	Costmin            string       `json:"costmin"`
	TickSize           string       `json:"tick_size"`
	Status             string       `json:"status"`
	LongPositionLimit  int          `json:"long_position_limit"`
	ShortPositionLimit int          `json:"short_position_limit"`
}

type Ticker struct {
	Ask             []string `json:"a"` //[<price>, <whole lot volume>, <lot volume>]
	Bid             []string `json:"b"` //[<price>, <whole lot volume>, <lot volume>]
	LastTradeClosed []string `json:"c"` //[<price>, <lot volume>]
	Volume          []string `json:"v"` //[<today>, <last 24 hours>]
	AveragePrice    []string `json:"p"` //[<today>, <last 24 hours>]
	Trades          []int    `json:"t"` //[<today>, <last 24 hours>]
	Low             []string `json:"l"` //[<today>, <last 24 hours>]
	High            []string `json:"h"` //[<today>, <last 24 hours>]
	Open            string   `json:"o"`
}

type GetOHCLVRequest struct {
	Pair     string `url:"pair"`
	Interval int    `url:"interval,omitempty"`
	Since    int64  `url:"since,omitempty"`
}

type GetOHCLVResponse struct {
	Last int      `json:"last"`
	Name []Candle `json:"name"`
}

type Candle struct {
	Timestamp int64
	Open      string
	High      string
	Low       string
	Close     string
	Vwap      string
	Volume    string
	Count     int
}
