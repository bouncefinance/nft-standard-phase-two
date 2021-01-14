package logger

import (
	"log"
	"nft_standard/config"
	"os"
	"testing"
)

func Test(t *testing.T)  {
	dir, _ := config.GetAppPath()
	logPath := dir + `\log\log.txt`

	_, err := os.Create(logPath)
	if err != nil {
		log.Fatal("create log file failed, error: ", err)
	}
}
