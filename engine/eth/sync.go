package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/logger"
	"nft_standard/models"
	"nft_standard/pkg/eth/geth"
	"nft_standard/pkg/gpool"
)

var (
	Pool *gpool.Pool
)

func init() {
	Pool = gpool.NewPool(20)
	go Pool.Run()
}

func Sync() {
	for i := int64(10000000); i <= 10010000; i++ {
		j := big.NewInt(i)
		task := gpool.NewTask(func() error {
			err := Do(j)
			if err != nil {
				logger.Logger.Error().Int64("block number", i).Msgf("Do error: %s", err)
			}
			return err
		})

		Pool.EntryChannel <- task
	}
}

func Do(blockNum *big.Int) error {
	logger.Logger.Info().Int64("blockNum", blockNum.Int64()).Msg("开始获取block相关数据")

	block, err := geth.GetBlock(blockNum)
	if err != nil {
		return errors.Wrap(err, "geth.GetBlock error")
	}

	difficulty, _ := big.NewFloat(0).SetInt(block.Difficulty()).Float64()
	blockM := models.EthBlock{
		Hash:       block.Hash().String()[2:],
		Number:     uint(block.Number().Uint64()),
		Difficulty: difficulty,
		GasLimit:   uint(block.GasLimit()),
		GasUsed:    uint(block.GasUsed()),
		Nonce:      uint(block.Nonce()),
		Miner:      block.Coinbase().String()[2:],
		Time:       uint(block.Time()),
		NTx:        uint(block.Transactions().Len()),
		StateRoot:  block.Root().String()[2:],
	}
	if blockNum.Int64() != 0 {
		blockM.ParentHash = block.ParentHash().String()[2:]
	}
	err = blockM.Refresh()
	if err != nil {
		return errors.Wrap(err, "blockM.Refresh error")
	}

	txes := block.Transactions()
	for _, tx := range txes {
		gasPrice, _ := big.NewFloat(0).SetInt(tx.GasPrice()).Float64()
		value, _ := big.NewFloat(0).SetInt(tx.Value()).Float64()
		//	将交易入库
		txM := models.EthTx{
			BlockNumber: uint(block.Number().Uint64()),
			BlockHash:   block.Hash().String()[2:],
			Gas:         uint(tx.Gas()),
			GasPrice:    gasPrice,
			Hash:        tx.Hash().String()[2:],
			Nonce:       uint(tx.Nonce()),
			Value:       value,
		}
		if tx.To() != nil{
			txM.To = tx.To().String()[2:]
		}
		if msg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(1))); err == nil {
			txM.From = msg.From().Hex()[2:]
		}
		err = txM.Refresh()
		if err != nil {
			logger.Logger.Error().Uint("block blockNum", blockM.Number).Str("tx hash", txM.Hash).Msgf("txM.Create error: %s", err)
			continue
		}
		block_tx := models.EthBlockTx{
			BlockId: blockM.Id,
			TxId:    txM.Id,
		}
		err = block_tx.Refresh()
		if err != nil {
			logger.Logger.Error().Uint("block blockNum", blockM.Number).Str("tx hash", txM.Hash).Msgf("block_tx.Create error: %s", err)
			continue
		}
	}
	return nil
}
