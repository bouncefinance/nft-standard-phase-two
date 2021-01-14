package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestSupportInterface(t *testing.T) {
	c := Contract721{
		Address: common.HexToAddress("0xC7e5e9434f4a71e6dB978bD65B4D61D3593e5f27"),
	}
	b, err := c.SupportInterface()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(b)
}

func TestGet721BaseURI(t *testing.T) {
	contract := "f6C3Aa70f29B64BA74dd6Abe6728cf8e190011b5"

	c := Contract721{
		Address: common.HexToAddress(contract),
	}
	uri, err := c.GetTokenIDUri(1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(uri)
}

/*
func TestGet1155BaseURI(t *testing.T) {
	contract := "0x7f15017506978517Db9eb0Abd39d12D86B2Af395"
	uri, err := Get1155BaseURI(contract)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(uri)
}
*/
