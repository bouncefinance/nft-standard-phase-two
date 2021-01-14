package models

import (
	"fmt"
	"testing"
)

func TestNftContractRecord721_Refresh(t *testing.T) {
	m := &NftContractRecord721{}
	m.BlockNumber = 7319539
	//m.BaseUri = "https://generic.url"
	m.ContractAddress = "0xbC7D9f0353A47EDb51AbF8419A27474e61Cb733f"
	m.Refresh()
}

func TestNftContractRecord721_GetBlockNum(t *testing.T) {
	m := &NftContractRecord721{}
	//m.BlockNumber = 7319539
	m.ContractAddress = "0x0a129D76Dbe2f678DDC478493dCABd0b40436114"
	m_, err := m.GetBlockNum()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m_)

	mm := &BtcBlock{
		Hash: "sssssssss",
	}
	err = mm.Create()
	if err != nil {
		t.Error(err)
		return
	}
}
