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
	"nft_standard/logger"
	"nft_standard/models"
	"nft_standard/pkg/eth/erpc"
	"nft_standard/pkg/eth/solidity"
	"strconv"
	"strings"
)

type Contract721 struct {
	Address common.Address
	DB      models.DBInterface
}

func (c *Contract721) SupportInterface() (bool, error) {
	var interSignBytes4 [4]byte
	interSignBytes, _ := hex.DecodeString(config.ERC721_INTERFACE_ID)
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

func (c *Contract721) BaseURI() (string, error) {
	data, err := config.ABI.Pack("baseURI")
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
func (c *Contract721) GetTokenIDUri(tokenID int64) (string, error) {
	data, err := config.ABI.Pack("tokenURI", big.NewInt(tokenID))
	if err != nil {
		return "", errors.Wrap(err, "BaseUri721ABI.Pack error")
	}
	data, err = erpc.EthCallRPC(c.Address.String(), "0x"+hex.EncodeToString(data))
	if err != nil {
		return "", errors.Wrap(err, "EthCallRPC error")
	}
	if len(data) == 0 {
		return "", config.ERROR_GET_TOKENID_URI_NULL
	}

	decodeString, err := solidity.ABIDecodeString(data)
	if err != nil {
		return "", err
	}
	return decodeString, nil
}

func (c *Contract721) SetDBBaseURI(s string) {
	c.DB.SetBaseURI(s)
}
func (c *Contract721) DBRefresh() error {
	return c.DB.Refresh()
}
func (c *Contract721) Has() (bool, error) {
	return c.DB.Has()
}
func (c *Contract721) GetBlockNum() (uint, error) {
	return c.DB.GetBlockNum()
}

func (c *Contract721) GetAddress() common.Address {
	return c.Address
}

func (c *Contract721) GetSectionLogs(fromBlock, toBlock *big.Int) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.Address},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    [][]common.Hash{{config.Topic_Transfer721}},
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

func (c *Contract721) ParseLogAndRefresh(logs []types.Log) {
	assets := make(map[float64]*models.NftTokenAssets721, 0)
	for _, lg := range logs {
		ok, _, to, tokenID := c.ParseLog(lg)
		if !ok {
			logger.Logger.Error().Str("contract address",lg.Address.String()).Str("tx hash",lg.TxHash.String()).
				Msg("Contract721 ParseLogAndRefresh c.ParseLog error: topic is not Transfer event")
			continue
		}
		assets[float64(tokenID)] = &models.NftTokenAssets721{
			ContractAddress: c.Address.String()[2:],
			TokenId:         float64(tokenID),
			OwnerAddress:    to[2:],
		}
	}

	for _, assets721 := range assets {
		err := c.RefreshAssets(assets721)
		if err != nil {
			logger.Logger.Error().Str("contract address",assets721.ContractAddress).Int("tokenID",int(assets721.TokenId)).
				Msgf("Contract721 ParseLogAndRefresh c.RefreshAssets error: %s",err)
			continue
		}
	}
	return
}

func (c *Contract721) RefreshAssets(m *models.NftTokenAssets721) error {
	has,m_, err := m.Has()
	if err != nil {
		return errors.Wrap(err, "m.Has error")
	}
	if !has {
		//	更新
		m.Id = m_.Id
		err := m.Update()
		if err != nil {
			return errors.Wrap(err, "m.Update error")
		}
		logger.Logger.Info().Interface("assets",m).Msg("更新一条 NftTokenAssets721 数据")
	} else {
		//	获取tokenURI
		uri, err := c.GetTokenIDUri(int64(m.TokenId))
		if err != nil {
			return errors.Wrap(err, "c.GetTokenIDUri error")
		}
		m.Uri = uri
		//	插入
		err = m.Create()
		if err != nil {
			return errors.Wrap(err, "m.Create error")
		}
		logger.Logger.Info().Interface("assets",m).Msg("创建一条 NftTokenAssets721 数据")
	}
	return nil
}

func (c *Contract721) ParseLog(log types.Log) (ok bool, from, to string, tokenID int) {
	var topics [4]string

	for i := range log.Topics {
		topics[i] = log.Topics[i].Hex()
	}
	if topics[0] != config.Topic_Transfer721.Hex() {
		return false, "", "", 0
	}

	from = "0x" + strings.Split(topics[1], "0x000000000000000000000000")[1]
	to = "0x" + strings.Split(topics[2], "0x000000000000000000000000")[1]
	tem, _ := strconv.ParseInt(topics[3], 0, 0)
	tokenID = int(tem)
	ok = true
	return
}
