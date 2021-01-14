package addr

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"nft_standard/config"
	"nft_standard/logger"
	"nft_standard/pkg/eth/erpc"
	"time"
)

func SubscribeBlock() {
	headers := make(chan *types.Header)
	sub, err := config.EthWssClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		logger.Logger.Error().Msgf("client.SubscribeNewHead() error: %s",err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			logger.Logger.Error().Msgf("sub.Err() error: %s",err)
			sub, err = config.EthWssClient.SubscribeNewHead(context.Background(), headers)
			if err != nil {
				logger.Logger.Error().Msgf("sub.Err() client.SubscribeNewHead() error: %s",err)
			}
			continue
		case header := <-headers:
			block, err := config.EthWssClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				logger.Logger.Error().Int64("block number",header.Number.Int64()).Msgf("client.BlockByHash() error: %s",err)
				continue
			}

			logger.Logger.Info().Int64("block number",header.Number.Int64()).Int("tx number",block.Transactions().Len()).Msg("监听到新block")
			go handleBlock(block)
		}
	}
}

func MonitorBlockOnce() {
	latest, err := erpc.GetBlockNum()
	if err != nil {
		logger.Logger.Error().Msgf("MonitorBlockOnce erpc.GetBlockNum() error: %s",err)
		return
	}
	time.Sleep(time.Second)

}










