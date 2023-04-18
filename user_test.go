package kraken

import (
	"fmt"
	"github.com/DawnKosmos/kraken-go/kapi"
	"testing"
)

func TestAccountBalance(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: "NgXwclr91aZZGoDr5u8+Zvct4JxKV5GzrxN81od/Dc1EbQTG3n9swuO4",
		SecretKey: "",
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
		PublicKey: "NgXwclr91aZZGoDr5u8+Zvct4JxKV5GzrxN81od/Dc1EbQTG3n9swuO4",
		SecretKey: "",
	}, true)

	resp, err := cl.AddOrder(kapi.AddOrderRequest{
		Pair:      "SOLUSD",
		Type:      "sell",
		Ordertype: "limit",
		Price:     "30.00",
		Volume:    "0.5",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(resp)

}
