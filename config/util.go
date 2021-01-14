package config

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	OPERATION_PARTICIPATE = 1
	EVENT_LOG_SECTION = 5000
	SYNC_LOG_GO_NUM = 30
)

func GetAppPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	index := strings.LastIndex(p, string(os.PathSeparator))
	return p[:index], nil
}

func GetConfigMsg(section, key string) string {
	s, _ := Cfg.GetSection(section)
	return s.Key(key).String()
}

