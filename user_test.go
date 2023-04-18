package kraken

import (
	"fmt"
	"github.com/DawnKosmos/kraken-go/kapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

const pk = "MJjgvvSJhB0mMj9Ps/Vv/BnvIkljPchRXKpTnWV+pgsP2nIuEa7fZ4qc"
const sk = ""

func TestAccountBalance(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: pk,
		SecretKey: sk,
	}, true)

	resp, err := cl.GetAccountBalance()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(resp)
}

func TestAddOrder(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: pk,
		SecretKey: sk,
	}, true)

	//Create Order
	resp, err := cl.AddOrder(kapi.AddOrderRequest{
		Pair:      "SOLUSD",
		Type:      "sell",
		Ordertype: "limit",
		Price:     "30.00",
		Volume:    "0.5",
		Leverage:  "4:1",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(resp)
	// Cancel Order
	resp2, err := cl.CancelOrder(resp.Txid[0])
	assert.Nil(t, err)
	fmt.Println(resp2)
}

func TestAddOrderTakeProfit(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: pk,
		SecretKey: sk,
	}, true)

	// TAKE PROFIT EXAMPLE
	resp, err := cl.AddOrderWithTakeProfit(kapi.AddOrderTakeProfit{
		Pair:            "SOLUSD",
		Type:            "sell",
		Ordertype:       "limit",
		Price:           "30.00",
		Volume:          "0.5",
		Leverage:        "4:1",
		TakeProfit:      "take-profit",
		TakeProfitPrice: "22.00",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(resp)

	resp2, err := cl.GetOpenOrders()
	assert.Nil(t, err)

	fmt.Println(resp2)

}
