package contract

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
	"nft_standard/config"
	"nft_standard/models"
	"nft_standard/pkg/eth/erpc"
	"nft_standard/pkg/eth/solidity"
)

type Contract1155 struct {
	Address common.Address
	DB      models.DBInterface
}

func (c *Contract1155) SupportInterface() (bool, error) {
	var interSignBytes4 [4]byte
	interSignBytes, _ := hex.DecodeString(config.ERC1155_INTERFACE_ID)
	for i := 0; i < len(interSignBytes4); i++ {
		interSignBytes4[i] = interSignBytes[i]
	}

	bytes, err := config.ABI.Pack("supportsInterface", interSignBytes4)
	if err != nil {
		return false, errors.Wrap(err, "arg.Pack error")
	}

	data, err := erpc.EthCallRPC(c.Address.String(), "0x"+hex.EncodeToString(bytes))
	if err != nil {
		return false, errors.Wrap(err, "EthCallRPC error")
	}
	if len(data) == 0 {
		return false, config.ERROR_SURPPORT_INTERFACE_NULL_165
	}

	b, err := solidity.ABIDecodeBool(data)
	if err != nil {
		return false, errors.Wrap(err, "solidity.ABIDecodeBool error")
	}
	return b, nil
}

func (c *Contract1155) BaseURI() (string, error) {
	data, err := config.ABI.Pack("uri", big.NewInt(0))
	if err != nil {
		return "", errors.Wrap(err, "BaseUri721ABI.Pack error")
	}
	data, err = erpc.EthCallRPC(c.Address.String(), "0x"+hex.EncodeToString(data))
	if err != nil {
		return "", errors.Wrap(err, "EthCallRPC error")
	}
	if len(data) == 0 {
		return "", config.ERROR_GET_BASE_URI_NULL
	}

	decodeString, err := solidity.ABIDecodeString(data)
	if err != nil {
		return "", err
	}
	return decodeString, nil
}

func (c *Contract1155) SetDBBaseURI(s string) {
	c.DB.SetBaseURI(s)
}
func (c *Contract1155) DBRefresh() error {
	return c.DB.Refresh()
}
func (c *Contract1155) Has() (bool, error) {
	return c.DB.Has()
}

//	字段相关
func (c *Contract1155)GetAddress()common.Address{
	return c.Address
}

func (c *Contract1155) GetBlockNum()(uint, error){
	return c.DB.GetBlockNum()
}

func (c *Contract1155)GetSectionLogs(fromBlock, toBlock *big.Int)([]types.Log, error){
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.Address},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    [][]common.Hash{{config.Topic_Transfer1155Single,config.Topic_Transfer1155Batch,config.Topic_Transfer1155URI}},
	}

	client, _, _ := config.ProjectIDS.GetMin()
	logs, err := client.HttpClient.FilterLogs(context.Background(), query)
	client.Decrease()

	time := 0
	for err != nil && time < config.RETRY_TIME {
		client, _, _ := config.ProjectIDS.GetMin()
		logs, err = client.HttpClient.FilterLogs(context.Background(), query)
		client.Decrease()
		time++
	}
	if err != nil {
		return logs, errors.Wrap(err, "client FilterLogs error")
	}
	return logs, nil
}
func (c *Contract1155) ParseLogAndRefresh(logs []types.Log) {}