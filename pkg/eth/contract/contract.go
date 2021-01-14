package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Contract interface {
	SupportInterface() (bool, error)
	BaseURI() (string, error)
	SetDBBaseURI(string)
	DBRefresh() error
	Has() (bool, error)
	GetAddress()common.Address
	GetBlockNum()(uint, error)
	GetSectionLogs(fromBlock, toBlock *big.Int)([]types.Log, error)
	ParseLogAndRefresh(logs []types.Log)
}


