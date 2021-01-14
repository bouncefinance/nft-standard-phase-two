package util

import "github.com/ethereum/go-ethereum/common"

type Param struct {
	//From     string `json:"from"`
	To string `json:"to"`
	//Gas      uint   `json:"gas"`
	//GasPrice uint   `json:"gasPrice"`
	//Value    uint   `json:"value"`
	Data string `json:"data"`
}
type RPCData struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      int    `json:"id"`
	Params  []interface{}
}
type BTCRPCData struct {
	JsonRPC   string        `json:"jsonrpc"`
	Method    string        `json:"method"`
	Id        string        `json:"id"`
	Params    []interface{} `json:"params"`
}
type RPCReturnData struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}
type BTCRPCReturnData struct {
	ErrorMsg interface{} `json:"error"`
	Id       string      `json:"id"`
	Result   interface{} `json:"result"`
}
type EtherscanReturnData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type SubData struct {
	Address common.Address
	Topics  []common.Hash
}

type EtherscanTXData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  []TXReceipt `json:"result"`
}

type TXReceipt struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

type RPCGetLogs struct {
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Address   string   `json:"address"`
	Topics    []string `json:"topics"`
}
