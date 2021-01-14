package logger

import (
	"github.com/rs/zerolog"
	"log"
	"nft_standard/config"
	"nft_standard/pkg/util"
	"os"
)

var Logger zerolog.Logger

func init() {
	logPath := config.Dir + config.LogPath
	file := &os.File{}
	var err error
	if util.FileNotExist(logPath) {
		file, err = os.Create(logPath)
		if err != nil {
			log.Fatal("create log file failed, error: ", err)
		}
	}else {
		file, err = os.OpenFile(logPath,os.O_RDWR,0666)
		if err != nil {
			log.Fatal("open log file failed, error: ", err)
		}
	}

	Logger = zerolog.New(file).With().Timestamp().Logger()
	Logger.Level(zerolog.InfoLevel)
}
