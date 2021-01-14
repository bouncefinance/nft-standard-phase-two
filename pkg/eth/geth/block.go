package geth

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/config"
)

func GetBlock(blockNum *big.Int) (*types.Block, error) {
	proIDFlag, _, _ := config.ProjectIDS.GetMin()
	block, err := proIDFlag.HttpClient.BlockByNumber(context.Background(), blockNum)
	proIDFlag.Decrease()

	time := 0
	for err != nil && time < config.RETRY_TIME {
		block, err = proIDFlag.HttpClient.BlockByNumber(context.Background(), blockNum)
		proIDFlag.Decrease()
		time++
	}
	if err != nil {
		return nil, errors.Wrap(err, "HttpClient.BlockByNumber error")
	}

	return block, nil
}
