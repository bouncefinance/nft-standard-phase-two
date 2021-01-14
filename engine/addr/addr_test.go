package addr

import (
	"math/big"
	"testing"
)

func TestGetAddrs(t *testing.T) {
	//bytes, err := erpc.GetBlockNum()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	end := big.NewInt(10918511)
	start := big.NewInt(10918511)
	//end.SetBytes(bytes)
	//start.SetBytes(bytes)
	//start.Sub(end,big.NewInt(5000))
	end.Add(end,big.NewInt(10000))	//	3600000
	GetContractAddrBaseURI(start, end)
}
