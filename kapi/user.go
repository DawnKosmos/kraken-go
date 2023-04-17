package kapi

type AddOrderRequest struct {
	Pair       string `url:"pair"`
	Type       string `url:"type"`
	Ordertype  string `url:"ordertype"`       // limit, market
	Price      string `url:"price,omitempty"` //
	Volume     string `url:"volume"`
	Leverage   string `url:"leverage,omitempty"`
	ReduceOnly bool   `url:"reduce_only,omitempty"`
}

type AddOrderResponse struct {
	Description map[string]any `json:"descr"`
	Txid        string         `json:"txid"`
}
