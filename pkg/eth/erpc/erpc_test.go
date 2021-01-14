package erpc

import (
	"fmt"
	"testing"
)

func TestSupportInterface(t *testing.T) {
	/*contractAddress := "0x93e508f373690cC4307a7A2363e573E63dAEF54E"
	supportInterface, err := contract.SupportInterface(contractAddress, config.ERC1155_INTERFACE_ID)
	if err != nil {
		t.Error(err)
		return
	}
	decodeBool, err := solidity.ABIDecodeBool(supportInterface)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(decodeBool)*/
}

func TestGetBlockNum(t *testing.T) {
	blockNum, err := GetBlockNum()
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(blockNum.Int64())
}



