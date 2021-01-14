package brpc

import (
	"encoding/json"
	"github.com/pkg/errors"
	"nft_standard/config"
	"nft_standard/pkg/util"
)

const (
	JSONRPC                    = "1.0"

	BTC_BLOCK_COUNT_METHOD     = "getblockcount"
	BTC_BLOCK_HASH_METHOD      = "getblockhash"
	BTC_BLOCK_METHOD           = "getblock"

	BTC_TRANSACTION_RAW_METHOD = "getrawtransaction"

	NodeURL = ""
)

func GetRawTransactionByHash(hash string)(map[string]interface{}, error) {
	result, err := brpcHelper(BTC_TRANSACTION_RAW_METHOD, "1", []interface{}{hash,true})
	if err != nil {
		return nil, errors.Wrap(err, "brpcHelper error")
	}

	tx, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("result type error")
	}
	return tx, nil
}

func GetBlockByHash(hash string) (map[string]interface{}, error) {
	result, err := brpcHelper(BTC_BLOCK_METHOD, "1", []interface{}{hash})
	if err != nil {
		return nil, errors.Wrap(err, "brpcHelper error")
	}

	block, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("result type error")
	}
	return block, nil
}

func GetBlockHashByNumber(number int) (string, error) {
	result, err := brpcHelper(BTC_BLOCK_HASH_METHOD, "1", []interface{}{number})
	if err != nil {
		return "", errors.Wrap(err, "brpcHelper error")
	}

	hash, ok := result.(string)
	if !ok {
		return "", errors.New("result type error")
	}
	return hash, nil
}

func GetBlockCount() (int, error) {
	result, err := brpcHelper(BTC_BLOCK_COUNT_METHOD, "1", nil)
	if err != nil {
		return -1, errors.Wrap(err, "brpcHelper error")
	}
	count, ok := result.(float64)
	if !ok {
		return -1, errors.New("result type error")
	}

	return int(count), nil
}

func brpcHelper(method, id string, params []interface{}) (interface{}, error) {
	rpcData := util.BTCRPCData{
		JsonRPC: JSONRPC,
		Method:  method,
		Id:      id,
		Params:  params,
	}
	data, err := util.PostUrlRetry(NodeURL, nil, rpcData, map[string]string{config.CONTENT_TYPE: config.CONTENT_TYPE_TEXT})
	if err != nil {
		return nil, errors.Wrap(err, "util.PostURL error")
	}

	result := util.BTCRPCReturnData{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal error")
	}
	if result.ErrorMsg != nil {
		return nil, errors.Wrap(errors.Errorf("%s", result.ErrorMsg), "result.ErrorMsg error")
	}

	return result.Result, err
}
