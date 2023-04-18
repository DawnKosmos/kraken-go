package kapi

type AddOrderRequest struct {
	Pair       string `url:"pair"`
	Type       string `url:"type"`            // Buy, Sell
	Ordertype  string `url:"ordertype"`       // limit, market
	Price      string `url:"price,omitempty"` //
	Volume     string `url:"volume"`
	Leverage   string `url:"leverage,omitempty"` // 2:1
	ReduceOnly bool   `url:"reduce_only,omitempty"`
}

type AddOrderResponse struct {
	Description map[string]any `json:"descr"`
	Txid        string         `json:"txid"`
}
