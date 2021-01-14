package main

import (
	"github.com/gin-gonic/gin"
	"math/big"
	"nft_standard/config"
	"nft_standard/engine/addr"
	"nft_standard/logger"
	"nft_standard/pkg/eth/erpc"
)

func main() {
	end, err := erpc.GetBlockNum()
	if err != nil {
		panic("erpc.GetBlockNum() error ===")
	}
	start := big.NewInt(0)

	go addr.GetContractAddrBaseURI(start, end)

	router := gin.Default()

	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error().Msgf("main recover: %s", r)
		}
	}()

	logger.Logger.Info().Str("port", config.Port).Msg("service is running")
	router.Run("0.0.0.0:" + config.AddressPort)
}
