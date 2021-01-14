package brpc

import (
	"fmt"
	"testing"
)

func TestGetBlockCount(t *testing.T) {
	count, err := GetBlockCount()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(count)
}

func TestGetBlockHashByNumber(t *testing.T) {
	/*count, err := GetBlockCount()
	if err != nil {
		t.Error(err)
		return
	}*/

	//count-=10000
	hash, err := GetBlockHashByNumber(0)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(hash)
}

func TestGetBlockByHash(t *testing.T) {
	count, err := GetBlockCount()
	if err != nil {
		t.Error(err)
		return
	}

	hash, err := GetBlockHashByNumber(count)
	if err != nil {
		t.Error(err)
		return
	}

	block, err := GetBlockByHash(hash)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(block)
}

func TestGetRawTransactionByHash(t *testing.T) {
	count, err := GetBlockCount()
	if err != nil {
		t.Error(err)
		return
	}

	hash, err := GetBlockHashByNumber(count)
	if err != nil {
		t.Error(err)
		return
	}

	block, err := GetBlockByHash(hash)
	if err != nil {
		t.Error(err)
		return
	}

	txes:= block["tx"].([]interface{})
	for _, tx_ := range txes {
		txHex:= tx_.(string)
		tx,err := GetRawTransactionByHash(txHex)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(tx)
	}
}