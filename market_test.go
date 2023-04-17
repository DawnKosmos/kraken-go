package kraken

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestMarket(t *testing.T) {
	cl, _ := New(nil, nil, true)

	res, err := cl.GetServerTime()
	assert.Nil(t, err)

	fmt.Println(res)

}
