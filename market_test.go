package kraken

import (
	"fmt"
	"github.com/DawnKosmos/kraken-go/kapi"
	"testing"
	"time"
)
import "github.com/stretchr/testify/assert"

func TestMarket(t *testing.T) {
	cl, _ := New(nil, nil, true)

	res4, err := cl.GetOHLC(kapi.GetOHLCRequest{
		Pair:     "USDC/USD",
		Interval: 1,
		Since:    time.Now().Add(-time.Hour * 1).Unix(),
	})
	assert.Nil(t, err)

	fmt.Println(res4.Pair)
	for _, v := range res4.Chart {
		fmt.Println(v)
	}
}
