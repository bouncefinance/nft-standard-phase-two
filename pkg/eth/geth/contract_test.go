package geth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestIsContract(t *testing.T) {
	contract := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

	b, err := IsContract(common.HexToAddress(contract))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(b)
}
