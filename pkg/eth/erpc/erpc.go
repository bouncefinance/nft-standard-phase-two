package erpc

import (
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/config"
	"nft_standard/pkg/util"
)

const (
	JSONRPC                 = "2.0"
	ETH_CALL_METHOD         = "eth_call"
	ETH_BLOCK_NUMBER_METHOD = "eth_blockNumber"
	ETH_GETLOGS_METHOD      = "eth_getLogs"
	BLOCK                   = "latest"
	request_id              = 1


)

func EthCallRPC(contractAddr string, dataStr string) ([]byte, error) {
	param := util.Param{
		To:   contractAddr,
		Data: dataStr,
	}

	rpcData := util.RPCData{
		JsonRPC: JSONRPC,
		Method:  ETH_CALL_METHOD,
		Id:      request_id,
		Params:  make([]interface{}, 0),
	}

	rpcData.Params = append(rpcData.Params, param)
	rpcData.Params = append(rpcData.Params, BLOCK)

	p, _, _ := config.ProjectIDS.GetMin()
	url := config.RPCHttpURL + p.PorID
	byte, err := util.PostUrlRetry(url, nil, rpcData, map[string]string{config.CONTENT_TYPE: config.CONTENT_TYPE_JSON})
	p.Decrease()
	if err != nil {
		return nil, errors.Wrap(err, "util.PostUrl error")
	}
	result := util.RPCReturnData{}
	err = json.Unmarshal(byte, &result)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal error")
	}
	if len(result.Result) == 0 {
		return nil, nil
	}
	if len(result.Result[2:]) == 0 {
		return nil, nil
	}
	bytes, err := hex.DecodeString(result.Result[2:])
	if err != nil {
		return nil, errors.Wrap(err, "hex.DecodeString error")
	}

	return bytes, nil
}

func GetBlockNum() (*big.Int, error) {
	rpcData := util.RPCData{
		JsonRPC: JSONRPC,
		Method:  ETH_BLOCK_NUMBER_METHOD,
		Id:      request_id,
		Params:  nil,
	}
	p, _, _ := config.ProjectIDS.GetMin()
	url := config.RPCHttpURL + p.PorID
	data, err := util.PostUrlRetry(url, nil, rpcData, map[string]string{config.CONTENT_TYPE: config.CONTENT_TYPE_JSON})
	p.Decrease()
	if err != nil {
		return nil, errors.Wrap(err, "util.PostURL error")
	}

	result := util.RPCReturnData{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal error")
	}
	if len(result.Result) == 0 {
		return nil, nil
	}
	bytes, err := hex.DecodeString(result.Result[2:])
	if err != nil {
		return nil, errors.Wrap(err, "hex.DecodeString error")
	}

	blockNum := big.NewInt(0)
	blockNum.SetBytes(bytes)

	return blockNum, nil
}

