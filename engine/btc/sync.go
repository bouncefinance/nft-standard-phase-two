package btc

import (
	"github.com/pkg/errors"
	"nft_standard/logger"
	"nft_standard/models"
	"nft_standard/pkg/btc/brpc"
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
	count, err := brpc.GetBlockCount()
	if err != nil {
		logger.Logger.Error().Msgf("brpc.GetBlockCount error: %s", err)
		return
	}
	for i := 0; i <= count; i++ {
		j := i
		task := gpool.NewTask(func() error {
			err := Do(j)
			if err != nil {
				logger.Logger.Error().Int("block height", i).Msgf("Do error: %s", err)
			}
			return err
		})

		Pool.EntryChannel <- task
	}
}

func Do(height int) error {
	logger.Logger.Info().Int("height", height).Msg("开始获取block相关数据")

	hash, err := brpc.GetBlockHashByNumber(height)
	if err != nil {
		return errors.Wrap(err, "brpc.GetBlockHashByNumber error")
	}
	block, err := brpc.GetBlockByHash(hash)
	if err != nil {
		return errors.Wrap(err, "brpc.GetBlockByHash error")
	}

	blockM := models.BtcBlock{
		Hash:          block["hash"].(string),
		Height:        uint(block["height"].(float64)),
		Confirmations: uint(block["confirmations"].(float64)),
		Nonce:         uint(block["nonce"].(float64)),
		Difficulty:    float32(block["difficulty"].(float64)),
		Size:          uint(block["size"].(float64)),
		StrippedSize:  uint(block["strippedsize"].(float64)),
		Time:          uint(block["time"].(float64)),
		NTx:           uint(block["nTx"].(float64)),
		MerkleRoot:    block["merkleroot"].(string),
		MedianTime:    uint(block["mediantime"].(float64)),
		ChainWork:     block["chainwork"].(string),
		Weight:        uint(block["weight"].(float64)),
		Version:       int(block["version"].(float64)),
		VersionHex:    block["versionHex"].(string),
		Bits:          block["bits"].(string),
	}
	if height != 0 {
		blockM.PreviousBlockHash = block["previousblockhash"].(string)
	}
	err = blockM.Refresh()
	if err != nil {
		return errors.Wrap(err, "blockM.Refresh error")
	}
	logger.Logger.Info().Interface("block", blockM).Msg("写入新区块数据")

	txes := block["tx"].([]interface{})
	for _, tx_ := range txes {
		txHex, ok := tx_.(string)
		if !ok {
			logger.Logger.Error().Uint("block height", blockM.Height).Interface("txHex", tx_).Msg("txHex type error")
			continue
		}
		tx, err := brpc.GetRawTransactionByHash(txHex)
		if err != nil {
			logger.Logger.Error().Uint("block height", blockM.Height).Str("tx_id", txHex).Msgf("brpc.GetRawTransactionByHash error: %s", err)
			continue
		}
		txM := models.BtcTx{
			Version:       uint(tx["version"].(float64)),
			VSize:         uint(tx["vsize"].(float64)),
			LockTime:      uint(tx["locktime"].(float64)),
			BlockHash:     tx["blockhash"].(string),
			TxId:          tx["txid"].(string),
			Hash:          tx["hash"].(string),
			Size:          uint(tx["size"].(float64)),
			Weight:        uint(tx["weight"].(float64)),
			Confirmations: uint(tx["confirmations"].(float64)),
			Time:          uint(tx["time"].(float64)),
			BlockTime:     uint(tx["blocktime"].(float64)),
		}
		err = txM.Refresh()
		if err != nil {
			logger.Logger.Error().Uint("block height", blockM.Height).Str("tx_id", txM.TxId).Msgf("txM.Create error: %s", err)
			continue
		}
		logger.Logger.Info().Interface("tx", txM).Msg("写入新交易数据")

		block_tx := models.BtcBlockTx{
			BlockId: blockM.Id,
			TxId:    txM.Id,
		}
		err = block_tx.Refresh()
		if err != nil {
			logger.Logger.Error().Uint("block height", blockM.Height).Str("tx_id", txM.TxId).Msgf("block_tx.Create error: %s", err)
			continue
		}
		logger.Logger.Info().Interface("block_tx", block_tx).Msg("写入区块交易中间表数据")
	}
	return nil
}
