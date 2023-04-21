package kraken

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const pk = "MJjgvvSJhB0mMj9Ps/Vv/BnvIkljPchRXKpTnWV+pgsP2nIuEa7fZ4qc"
const sk = "9uyh9XWVCdhbjO3WEp7i3uOlIV1bRfmB25iPXw2M+dCP9R6hoScFAfvEqL/7c5UB37dKmiyOIrgUk61Hk09mRQ=="

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
	/*
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
	*/
}

func TestAddOrderTakeProfit(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: pk,
		SecretKey: sk,
	}, true)

	resp2, err := cl.GetOpenOrders()
	assert.Nil(t, err)

	for k, v := range resp2.Open {
		fmt.Println(k, v)
	}
}
