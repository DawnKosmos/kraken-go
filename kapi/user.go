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

type AddOrderTakeProfit struct {
	Pair            string `url:"pair"`
	Type            string `url:"type"`            // Buy, Sell
	Ordertype       string `url:"ordertype"`       // limit, market
	Price           string `url:"price,omitempty"` //
	Volume          string `url:"volume"`
	Leverage        string `url:"leverage,omitempty"` // 2:1
	ReduceOnly      bool   `url:"reduce_only,omitempty"`
	TakeProfit      string `url:"close[ordertype], omitempty"`
	TakeProfitPrice string `url:"close[price],omitempty"`
}

type AddOrderResponse struct {
	Description map[string]any `json:"descr"`
	Txid        []string       `json:"txid"`
}

type CancelOrderRequest struct {
	Txid string `url:"txid"`
}

type CancelOrderResponse struct {
	Count   int  `json:"count"`
	Pending bool `json:"pending"`
}

type OpenOrderResponse struct {
	Open map[string]any `url:"open"`
}
