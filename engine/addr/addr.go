package addr

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/config"
	"nft_standard/logger"
	"nft_standard/models"
	"nft_standard/pkg/eth/contract"
	"nft_standard/pkg/eth/geth"
	"nft_standard/pkg/gpool"
)

func GetContractAddrBaseURI(start, end *big.Int) {
	for i := start.Int64(); i < end.Int64(); i++ {
		number := big.NewInt(i)

		task := gpool.NewTask(func() error {
			block, err := geth.GetBlock(number)
			if err != nil {
				logger.Logger.Error().Int64("block number", i).Msgf("geth.GetBlock error: %s", err)
				return err
			}
			logger.Logger.Info().Uint64("block number", block.NumberU64()).Int("交易数量", block.Transactions().Len()).
				Str("block hash", block.Hash().String()).Msg("获取到block")
			handleBlock(block)
			return nil
		})

		gpool.PoolGO.EntryChannel <- task
	}
}

func recordMachine(contract_ contract.Contract) (bool, error) {
	b, err := contract_.SupportInterface()
	if err != nil && err != config.ERROR_SURPPORT_INTERFACE_NULL_165 {
		return false, errors.Wrap(err, "contract.SupportInterface error")
	}
	if b {
		baseURI, err := contract_.BaseURI()
		if err != nil && err != config.ERROR_GET_BASE_URI_NULL {
			logger.Logger.Error().Str("contract address", contract_.GetAddress().String()).
				Msgf("获取合约baseURI失败, error: %s", err)
		}
		contract_.SetDBBaseURI(baseURI)
		err = contract_.DBRefresh()
		if err != nil {
			return false, errors.Wrap(err, "contract.DBRefresh error")
		}
	}
	return b, nil
}

func handleBlock(block *types.Block,isSub bool) {
	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}
		m := &models.NftContractRecord{
			ContractAddress: tx.To().String(),
		}
		has, err := m.Has()
		if has {
			continue
		}
		if err != nil {
			logger.Logger.Error().Str("contract address", tx.To().String()).
				Int64("block number", block.Number().Int64()).
				Msgf("GetContractAddrBaseURI NftContractRecord Has() error: %s", err)
		}
		err = m.Create()
		if err != nil {
			logger.Logger.Error().Str("contract address", tx.To().String()).
				Int64("block number", block.Number().Int64()).
				Msgf("GetContractAddrBaseURI NftContractRecord Create() error: %s", err)
		}
		logger.Logger.Info().Str("contract address", tx.To().String()).
			Int64("block number", block.Number().Int64()).Msg("记录一条已判断过的地址")

		c721 := &contract.Contract721{
			Address: *tx.To(),
			DB: &models.NftContractRecord721{
				ContractAddress: tx.To().String(),
				BlockNumber:     uint(block.NumberU64()),
			},
		}
		c1155 := &contract.Contract1155{
			Address: *tx.To(),
			DB: &models.NftContractRecord1155{
				ContractAddress: tx.To().String(),
				BlockNumber:     uint(block.NumberU64()),
			},
		}
		isContract, err := geth.IsContract(*tx.To())
		if err != nil {
			logger.Logger.Error().Str("contract address", tx.To().String()).
				Int64("block number", block.Number().Int64()).
				Msgf("GetContractAddrBaseURI 721 geth.IsContract error: %s", err)
			continue
		}
		if isContract {
			b, err := recordMachine(c721)
			if err != nil {
				logger.Logger.Error().Str("contract address", tx.To().String()).
					Int64("block number", block.Number().Int64()).
					Msgf("GetContractAddrBaseURI 721 recordMachine error: %s", err)
			}
			if !b {
				b, err := recordMachine(c1155)
				if err != nil {
					logger.Logger.Error().Str("contract address", tx.To().String()).
						Int64("block number", block.Number().Int64()).
						Msgf("GetContractAddrBaseURI 1155 recordMachine error: %s", err)
					continue
				}
				if b {
					logger.Logger.Info().Str("contract address", tx.To().String()).
						Int64("block number", block.Number().Int64()).Msg("find a 1155 contract")
				}
			} else {
				logger.Logger.Info().Str("contract address", tx.To().String()).
					Int64("block number", block.Number().Int64()).Msg("find a 721 contract")
			}
		}
	}
	if isSub{

	}else {

	}
}
