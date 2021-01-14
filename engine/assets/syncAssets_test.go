package assets

import (
	"github.com/ethereum/go-ethereum/common"
	"nft_standard/models"
	"nft_standard/pkg/eth/contract"
	"testing"
)

func TestSyncHisAssets(t *testing.T) {
	contractAddr:="0xf6C3Aa70f29B64BA74dd6Abe6728cf8e190011b5"
	c:=&contract.Contract721{
		Address: common.HexToAddress(contractAddr),
		DB:      &models.NftContractRecord721{
			ContractAddress: contractAddr,
		},
	}

	SyncHisAssets(c)
}
