package kraken

import (
	"github.com/DawnKosmos/kraken-go/kapi"
	"strings"
)

func (c *Client) GetServerTime() (kapi.GetServerTimeResult, error) {
	var resp kapi.Response[kapi.GetServerTimeResult]
	err := c.GET("/0/public/Time", nil, &resp)
	return resp.Result, err
}

func (c *Client) GetAssetInfo(assets ...string) (map[string]kapi.Asset, error) {
	var resp kapi.Response[map[string]kapi.Asset]
	var err error
	if len(assets) == 0 {
		err = c.GET("/0/public/Assets", nil, &resp)
	} else {
		err = c.GET("/0/public/Assets", kapi.GetAssetInfoRequest{Asset: SumStringArray(assets)}, &resp)

	}
	return resp.Result, err
}

func (c *Client) GetTradableAssetPairs(pairs ...string) (map[string]kapi.TradableAsset, error) {
	var resp kapi.Response[map[string]kapi.TradableAsset]
	var err error
	if len(pairs) == 0 {
		err = c.GET("/0/public/AssetPairs", nil, &resp)
	} else {
		err = c.GET("/0/public/AssetPairs", kapi.GetTradablePairRequest{Pair: SumStringArray(pairs)}, &resp)
	}
	return resp.Result, err
}

func (c *Client) GetTicker(tickers ...string) (map[string]kapi.Ticker, error) {
	var resp kapi.Response[map[string]kapi.Ticker]
	var err error
	if len(tickers) == 0 {
		err = c.GET("/0/public/Ticker", nil, &resp)
	} else {
		err = c.GET("/0/public/Ticker", kapi.GetTradablePairRequest{Pair: SumStringArray(tickers)}, &resp)
	}
	return resp.Result, err
}

// GETOHCLV supports following intervals 1 5 15 30 60 240 1440 10080 21600
func (c *Client) GetOHLC(req kapi.GetOHCLVRequest) ([]byte, error) {
	var resp kapi.Response[[]byte]

	err := c.GET("/0/public/OHLC", req, &resp)
	return resp.Result, err

	// Try jsoniter
}

//==== HELPER FUNCTIONS

func SumStringArray(arr []string) string {
	var builder strings.Builder
	builder.WriteString(arr[0])
	for _, v := range arr[1:] {
		builder.WriteString(",")
		builder.WriteString(v)
	}
	return builder.String()
}
