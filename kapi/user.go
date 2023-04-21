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
	Open map[string]Order `url:"open"`
}

type Order struct {
	Refid    string  `json:"refid"`
	Userref  int     `json:"userref"`
	Status   string  `json:"status"`
	Opentm   float64 `json:"opentm"`
	Starttm  int     `json:"starttm"`
	Expiretm int     `json:"expiretm"`
	Descr    struct {
		Pair      string `json:"pair"`
		Type      string `json:"type"`
		Ordertype string `json:"ordertype"`
		Price     string `json:"price"`
		Price2    string `json:"price2"`
		Leverage  string `json:"leverage"`
		Order     string `json:"order"`
		Close     string `json:"close"`
	} `json:"descr"`
	Vol        string `json:"vol"`
	VolExec    string `json:"vol_exec"`
	Cost       string `json:"cost"`
	Fee        string `json:"fee"`
	Price      string `json:"price"`
	Stopprice  string `json:"stopprice"`
	Limitprice string `json:"limitprice"`
	Misc       string `json:"misc"`
	Oflags     string `json:"oflags"`
	Trades     string `json:"trades,omitempty"`
}
