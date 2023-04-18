package kraken

import "github.com/DawnKosmos/kraken-go/kapi"

func (c *Client) GetAccountBalance() (map[string]string, error) {
	var resp kapi.Response[map[string]string]
	err := c.POST("/0/private/Balance", nil, &resp)
	return resp.Result, err
}

func (c *Client) AddOrder(req kapi.AddOrderRequest) (kapi.AddOrderResponse, error) {
	var resp kapi.Response[kapi.AddOrderResponse]
	err := c.POST("/0/private/AddOrder", req, &resp)
	return resp.Result, err
}

func (c *Client) AddOrderWithTakeProfit(req kapi.AddOrderTakeProfit) (kapi.AddOrderResponse, error) {
	var resp kapi.Response[kapi.AddOrderResponse]
	err := c.POST("/0/private/AddOrder", req, &resp)
	return resp.Result, err
}

func (c *Client) CancelOrder(TxId string) (kapi.CancelOrderResponse, error) {
	var resp kapi.Response[kapi.CancelOrderResponse]
	err := c.POST("/0/private/CancelOrder", kapi.CancelOrderRequest{Txid: TxId}, &resp)
	return resp.Result, err
}

func (c *Client) GetOpenOrders() (kapi.OpenOrderResponse, error) {
	var resp kapi.Response[kapi.OpenOrderResponse]
	err := c.POST("/0/private/OpenOrders", nil, &resp)
	return resp.Result, err
}
