package config

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

const (
	//	node
	NODE_NUM = 1
	//	HTTP
	RETRY_TIME        = 5
	CONTENT_TYPE      = "Content-type"
	CONTENT_TYPE_JSON = "application/json"
	CONTENT_TYPE_TEXT = "text/plain"

	configPath = "config/config.ini"

	INT_MAX = int(^uint(0) >> 1)

	//	contract
	ERC721_INTERFACE_ID  = "80ac58cd"
	ERC1155_INTERFACE_ID = "d9b67a26"

	SafeTransferFrom721_01    = "0x42842e0e"
	SafeTransferFrom721_02    = "0xb88d4fde"
	TransferFrom721           = "0x23b872dd"
	SafeTransferFrom1155      = "0xf242432a"
	SafeBatchTransferFrom1155 = "0x2eb2c2d6"

	SUPPORT_INTERFACE_SIGN = "0x01ffc9a7" //	supportsInterface(bytes4 interfaceId)
	BaseURI721             = "0x6c0360eb" //	crypto.Keccak256Hash([]byte("baseURI()"))
	BaseURI1155            = "0x0e89341c" //	crypto.Keccak256Hash([]byte("uri(uint256)"))

	ABIPath = "config/abi.json"

	MongoURL = ""
)

//	mongo
const (
	DATABASE                 = "AnkrNFT"
	COLLECTION               = "LatestBlockNum"
	BASE_URI_1155_COLLECTION = "BaseURI1155"
	ADDRESS_NFT_COLLECTION   = "AddressNFT"
)

var (
	Dir string
	Cfg *ini.File

	IsBSC     = false
	IsMain    = false
	IsRinkeby = false

	//	HTTP
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RPCHttpURL   string
	ProjectIDS   *ProIDFlags
	AddressPort  string

	//	contract
	AddressZero              = common.Address{}
	ABI                      = abi.ABI{}
	BaseUri1155ABI           = abi.ABI{}
	Topic_Transfer721        = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	Topic_Transfer1155Single = crypto.Keccak256Hash([]byte("TransferSingle(address,address,address,uint256,uint256)"))
	Topic_Transfer1155Batch  = crypto.Keccak256Hash([]byte("TransferBatch(address,address,address,uint256[],uint256[])"))
	Topic_Transfer1155URI    = crypto.Keccak256Hash([]byte("URI(string,uint256)"))

	//	config
	DBSection        string
	ChainInfoSection string
	LogPath          string
	URLPath          string
	EthWssClient     *ethclient.Client
)

var (
	mode    string
	network string
)

func init() {
	var err error
	Dir, err = GetAppPath()
	if err != nil {
		panic(err)
	}

	Cfg, err = ini.Load(Dir + "/" + configPath)
	if err != nil {
		panic(err)
	}

	loadBase()
	initABI()
	ParseFlag()
}
func ParseFlag() {
	flag.StringVar(&mode, "m", "api", "指定运行模式，api 提供api服务 address 获取账户地址 sync 同步历史资产 subscribe 订阅新资产，并监控漏掉的数据")
	flag.StringVar(&network, "n", "main", "指定支持网络，支持三种网络 main rinkeby bsc")
	flag.Parse()

	initCommon()

	switch mode {
	case "api":
	case "address":
	case "sync":
	case "subscribe":
	}

	switch network {
	case "main":
		initMain()
	case "rinkeby":
		initRinkeby()
	case "bsc":
		initBsc()
	case "local":
		initBsc()
		DBSection = "databaseLocal"
	}

	logTime := time.Now().Unix()
	LogPath = fmt.Sprintf("/logger/%s/log_%s_%s_%d.log", mode, mode, network, logTime)
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func initABI() {
	s, err := os.OpenFile(Dir+"/"+ABIPath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	ABI, err = abi.JSON(s)
	if err != nil {
		panic(err)
	}
}

func initCommon() {
	server, err := Cfg.GetSection("server")
	if err != nil {
		panic("config get server section failed")
	}
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func initMain() {
	RPCHttpURL = GetConfigMsg("main", "RPC_URL")
	Port = GetConfigMsg("main", "HTTP_PORT")
	URLPath = "query_main"
	DBSection = "databaseMain"
	AddressPort = "9999"
	IsMain = true

	initHttpClient(1, nil)

	var err error
	EthWssClient, err = ethclient.Dial(GetConfigMsg("main", "WSS_URL") + GetConfigMsg("infura", "PROJECT_ID_1"))
	if err != nil {
		panic(err)
	}
}
func initRinkeby() {
	RPCHttpURL = GetConfigMsg("rinkeby", "RPC_URL")
	Port = GetConfigMsg("rinkeby", "HTTP_PORT")
	URLPath = "query_rinkeby"
	DBSection = "databaseRinkeby"
	AddressPort = "9998"
	IsRinkeby = true

	infura, err := Cfg.GetSection("infura")
	if err != nil {
		panic("config get infura section failed")
	}
	initHttpClient(30, infura)

	EthWssClient, err = ethclient.Dial(GetConfigMsg("rinkeby", "WSS_URL") + GetConfigMsg("infura", "PROJECT_ID_2"))
	if err != nil {
		panic(err)
	}
}
func initBsc() {
	RPCHttpURL = GetConfigMsg("bsc", "RPC_URL")
	Port = GetConfigMsg("bsc", "HTTP_PORT")
	URLPath = "query_bsc"
	DBSection = "databaseBsc"
	AddressPort = "9997"
	IsBSC = true

	initHttpClient(1, nil)

	var err error
	EthWssClient, err = ethclient.Dial(GetConfigMsg("bsc", "WSS_URL") + GetConfigMsg("infura", "PROJECT_ID_3"))
	if err != nil {
		panic(err)
	}
	go func() {
		ticker := time.NewTicker(time.Second * 40)
		for {
			select {
			case <-ticker.C:
				_, _ = EthWssClient.ChainID(context.Background())
			}
		}
	}()
}

func initHttpClient(number int, sec *ini.Section) {
	arr := make([]*ProIDFlag, NODE_NUM)
	for i := 0; i < number; i++ {
		if sec != nil {
			arr[i] = NewProIDFlag(sec.Key(fmt.Sprint("PROJECT_ID_", i+1)).String(), 0)
		} else {
			arr[i] = NewProIDFlag("", 0)
		}
	}
	ProjectIDS = NewProIDFlags(arr)
}