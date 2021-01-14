package assets

import (
	"math/big"
	"nft_standard/config"
	"nft_standard/logger"
	"nft_standard/pkg/eth/contract"
	"nft_standard/pkg/eth/erpc"
	"nft_standard/pkg/gpool"
	"nft_standard/pkg/util"
)

func SyncHisAssets(cont contract.Contract) {
	latestNum, err := erpc.GetBlockNum()
	if err != nil {
		logger.Logger.Error().Str("contract address",cont.GetAddress().String()).Msgf("SyncHisAssets erpc.GetBlockNum error: %s", err)
		return
	}

	startNum, err := cont.GetBlockNum()
	if err != nil {
		logger.Logger.Error().Str("contract address",cont.GetAddress().String()).Msgf("SyncHisAssets cont.GetBlockNum error: %s", err)
		return
	}

	time := util.DivideTime(int64(startNum), latestNum.Int64())

	logger.Logger.Info().Str("contract address",cont.GetAddress().String()).Uint("start",startNum).Int64("end",latestNum.Int64()).
		Int64("区间数",time).Msg("开始获取合约的所有历史日志")
	for i := int64(0); i < time; i++ {
		start := int64(startNum) + i*config.EVENT_LOG_SECTION
		end := int64(startNum) + (i+1)*config.EVENT_LOG_SECTION - 1
		startBig := big.NewInt(start)
		endBig := big.NewInt(end)
		if i == time-1 {
			endBig = latestNum
		}

		task := gpool.NewTask(func() error {
			logs, err := cont.GetSectionLogs(startBig, endBig)
			if err != nil {
				logger.Logger.Error().Str("contract address",cont.GetAddress().String()).Int64("start",startBig.Int64()).
					Int64("end",endBig.Int64()).Msgf("SyncHisAssets cont.GetSectionLogs error: %s", err)
				return err
			}
			logger.Logger.Info().Str("contract address",cont.GetAddress().String()).Int64("start",startBig.Int64()).
				Int64("end",endBig.Int64()).
				Int("日志数量",len(logs)).Msg("获取到合约的其中一个区间的历史日志")
			cont.ParseLogAndRefresh(logs)
			return nil
		})
		gpool.PoolGO.EntryChannel <- task
	}
}
