package geth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/config"
)

func IsContract(address common.Address) (bool, error) {
	client, _, _ := config.ProjectIDS.GetMin()
	byteCode, err := client.HttpClient.CodeAt(context.Background(), address, nil) // nil is latest block
	client.Decrease()
	if err != nil {
		return false, errors.Wrap(err,"client.codeAt() error")
	}

	return len(byteCode) > 0,nil
}

func SyncContractLog(contractAddr common.Address, eventSigns []common.Hash, fromBlock, toBlock *big.Int) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    [][]common.Hash{eventSigns},
	}

	client, _, _ := config.ProjectIDS.GetMin()
	logs, err := client.HttpClient.FilterLogs(context.Background(), query)
	client.Decrease()

	time := 0
	for err != nil && time < config.RETRY_TIME {
		client, _, _ := config.ProjectIDS.GetMin()
		logs, err = client.HttpClient.FilterLogs(context.Background(), query)
		client.Decrease()
		time++
	}
	if err != nil {
		return logs, errors.Wrap(err, "client FilterLogs error")
	}
	return logs, nil
}