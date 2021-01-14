package main

import (
	"fmt"
	"nft_standard/logger"
	"net/http"
	"nft_standard/config"
	"nft_standard/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}


	logger.Logger.Info().Str("port",config.Port).Msgf("service run at %s",config.Port)
	//router.RunTLS(":443",conf.Dir + "/" + "cert/4618025_nftview.bounce.finance.pem",conf.Dir + "/" + "cert/4618025_nftview.bounce.finance.key")
	s.ListenAndServe()
}

