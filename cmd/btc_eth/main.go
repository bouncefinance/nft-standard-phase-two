package main

import (
	"github.com/gin-gonic/gin"
	"nft_standard/engine/btc"
	"nft_standard/config"
	"nft_standard/logger"
)

func main() {
	go btc.Sync()
	//go eth.Sync()

	router := gin.Default()

	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Error().Msgf("main recover: %s", r)
		}
	}()

	logger.Logger.Info().Str("port", config.Port).Msg("service is running")
	router.Run("0.0.0.0:" + config.AddressPort)
}
