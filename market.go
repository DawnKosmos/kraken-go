package kraken

import (
	"github.com/DawnKosmos/kraken-go/kapi"
	"strconv"
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
func (c *Client) GetOHLC(req kapi.GetOHLCRequest) (kapi.GetOHLCResponse, error) {
	var resp kapi.Response[map[string]any]
	err := c.GET("/0/public/OHLC", req, &resp)
	if err != nil {
		return kapi.GetOHLCResponse{}, err
	}
	return convertToOHCL(resp.Result), nil
}

//==== HELPER FUNCTIONS

func convertToOHCL(resp map[string]any) kapi.GetOHLCResponse {
	var cc kapi.GetOHLCResponse
	for key, val := range resp {
		if key == "last" {
			cc.Last = int(val.(float64))
			continue
		}
		cc.Pair = key
		items := val.([]any)
		cc.Chart = make([]kapi.Candle, 0, len(items))
		var err error
		for _, v := range items {
			var temp kapi.Candle
			candle := v.([]interface{})
			temp.Timestamp = int64(candle[0].(float64))
			if temp.Open, err = strconv.ParseFloat(candle[1].(string), 64); err != nil {
				continue
			}
			if temp.High, err = strconv.ParseFloat(candle[2].(string), 64); err != nil {
				continue
			}
			if temp.Low, err = strconv.ParseFloat(candle[3].(string), 64); err != nil {
				continue
			}
			if temp.Close, err = strconv.ParseFloat(candle[4].(string), 64); err != nil {
				continue
			}
			if temp.Vwap, err = strconv.ParseFloat(candle[5].(string), 64); err != nil {
				continue
			}
			if temp.Volume, err = strconv.ParseFloat(candle[6].(string), 64); err != nil {
				continue
			}
			cc.Chart = append(cc.Chart, temp)
		}
	}
	return cc
}

func SumStringArray(arr []string) string {
	var builder strings.Builder
	builder.WriteString(arr[0])
	for _, v := range arr[1:] {
		builder.WriteString(",")
		builder.WriteString(v)
	}
	return builder.String()
}
