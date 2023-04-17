package kraken

import (
	"fmt"
	"testing"
)

func TestAccountBalance(t *testing.T) {
	cl, err := New(nil, &Account{
		PublicKey: "",
		SecretKey: "",
	}, true)

	resp, err := cl.GetAccountBalance()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(resp)
}
