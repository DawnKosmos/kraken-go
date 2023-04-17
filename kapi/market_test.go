package kapi

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMarket(t *testing.T) {
	var call []byte = []byte(`{"error":[],"result":{"XXBTZUSD":[[1680912000,"27923.0","28170.0","27886.0","27956.0","28017.9","814.32026236",15890],[1680998400,"27955.9","28548.7","27812.2","28333.7","28203.8","1685.88383105",21294],[1681084800,"28333.8","29800.0","28170.0","29655.0","28945.8","5357.58172588",46175],[1681171200,"29654.9","30586.0","29615.0","30227.6","30115.2","5638.76555055",53246],[1681257600,"30227.7","30505.5","29645.8","29895.2","30057.5","4582.12074125",42996],[1681344000,"29895.2","30650.8","29873.9","30405.9","30303.3","3994.17529269",38636],[1681430400,"30406.0","31050.0","30000.0","30493.2","30611.9","5123.14673656",47027],[1681516800,"30493.2","30624.1","30250.0","30315.9","30379.7","1579.39907208",19921],[1681603200,"30315.9","30555.6","30145.0","30316.0","30341.6","1432.96396029",18436],[1681689600,"30316.1","30319.3","29777.0","29889.7","29947.6","1431.51497736",13735]],"last":1681603200}}`)
	var resp Response[map[string]any]

	err := json.Unmarshal(call, &resp)
	assert.Nil(t, err)

	var cc GetOHLCResponse
	for key, val := range resp.Result {
		if key == "last" {
			cc.Last = int(val.(float64))
			continue
		}
		items := val.([]any)
		cc.Chart = make([]Candle, 0, len(items))
		for _, v := range items {
			var temp Candle
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

}
