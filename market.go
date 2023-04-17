package kraken

import "github.com/DawnKosmos/kraken-go/kapi"

func (c *Client) GetServerTime() (kapi.GetServerTimeResult, error) {
	var resp kapi.Response[kapi.GetServerTimeResult]
	err := c.GET("/0/public/Time", nil, &resp)
	return resp.Result, err
}
