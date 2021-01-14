package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	LOG_PATH = "../log1609515113.txt"
)

type GetErc struct {
	UserAddr     string `json:"userAddr"`
	ContractAddr string `json:"contractAddr"`
	ChainID      int    `json:"chainID"`
	Message      string `json:"message"`
}

type GetAllNFT struct {
	UserAddr string `json:"userAddr"`
	ChainID  int    `json:"chainID"`
	Message  string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

func main() {
	file, err := os.OpenFile(LOG_PATH, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("openFile failed error:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	Contain(r, "error","./log_deal.txt")
	//JsonMar(r)
}

func OpenLogFile() *bufio.Reader {
	file, err := os.OpenFile(LOG_PATH, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("openFile failed error:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	return r
}

func Contain(reader *bufio.Reader,subStr string,dealedPath string) {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}

		if strings.Contains(string(buf), subStr) {
			_, err := requestLog.Write(buf)
			if err != nil {
				log.Fatal("write file failed", err)
			}
		}
	}
}

func DisContain(reader *bufio.Reader,subStr string,dealedPath string) {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}

		if !strings.Contains(string(buf), subStr) {
			_, err := requestLog.Write(buf)
			if err != nil {
				log.Fatal("write file failed", err)
			}
		}
	}
}

func JsonMar(reader *bufio.Reader,dealedPath string)  {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	msg := Message{}
	i:=0
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
		err = json.Unmarshal(buf, &msg)
		if err!=nil{
			fmt.Println("json.Unmarshal err: ",err)
		}
		fmt.Println(msg.Message)
		i++
		if i == 1{
			break
		}
	}
}