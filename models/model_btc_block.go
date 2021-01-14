package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//BtcBlock
type BtcBlock struct {
	Id                uint    `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20) unsigned,PRI"`
	Hash              string  `gorm:"column:hash" form:"hash" json:"hash" comment:"区块Hash" sql:"varchar(64),UNI"`
	Height            uint    `gorm:"column:height" form:"height" json:"height" comment:"区块高度" sql:"int(10) unsigned,UNI"`
	Confirmations     uint    `gorm:"column:confirmations" form:"confirmations" json:"confirmations" comment:"区块确认数" sql:"int(10) unsigned"`
	Nonce             uint    `gorm:"column:nonce" form:"nonce" json:"nonce" comment:"nonce值" sql:"int(10) unsigned"`
	Difficulty        float32 `gorm:"column:difficulty" form:"difficulty" json:"difficulty" comment:"区块难度" sql:"decimal(32,2)"`
	Size              uint    `gorm:"column:size" form:"size" json:"size" comment:"区块字节数" sql:"int(10) unsigned"`
	StrippedSize      uint    `gorm:"column:stripped_size" form:"stripped_size" json:"stripped_size" comment:"剔除隔离见证数据后的区块字节数" sql:"int(10) unsigned"`
	Time              uint    `gorm:"column:time" form:"time" json:"time" comment:"区块创建时间戳" sql:"int(10) unsigned"`
	NTx               uint    `gorm:"column:n_tx" form:"n_tx" json:"n_tx" comment:"block包含的交易数量" sql:"int(10) unsigned"`
	MerkleRoot        string  `gorm:"column:merkle_root" form:"merkle_root" json:"merkle_root" comment:"区块的默克尔树根" sql:"varchar(64)"`
	MedianTime        uint    `gorm:"column:median_time" form:"median_time" json:"median_time" comment:"区块中值时间戳" sql:"int(10) unsigned"`
	ChainWork         string  `gorm:"column:chain_work" form:"chain_work" json:"chain_work" comment:"chainwork" sql:"varchar(64)"`
	PreviousBlockHash string  `gorm:"column:previous_block_hash" form:"previous_block_hash" json:"previous_block_hash" comment:"前一区块的哈希" sql:"varchar(64)"`
	Weight            uint    `gorm:"column:weight" form:"weight" json:"weight" comment:"BIP141定义的区块权重" sql:"int(10) unsigned"`
	Version           int     `gorm:"column:version" form:"version" json:"version" comment:"版本" sql:"int(11)"`
	VersionHex        string  `gorm:"column:version_hex" form:"version_hex" json:"version_hex" comment:"16进制表示的版本" sql:"varchar(24)"`
	Bits              string  `gorm:"column:bits" form:"bits" json:"bits" comment:"bits" sql:"varchar(24)"`
}

//TableName
func (m *BtcBlock) TableName() string {
	return "btc_block"
}

//One
func (m *BtcBlock) One() (one *BtcBlock, err error) {
	one = &BtcBlock{}
	err = crudOne(m, one)
	return
}

//All
func (m *BtcBlock) All(q *PaginationQuery) (list *[]BtcBlock, total uint, err error) {
	list = &[]BtcBlock{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *BtcBlock) Update() (err error) {
	where := BtcBlock{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *BtcBlock) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *BtcBlock) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

/**
 * @parameter:
 * @return:
 * @Description: 判断数据库中是否存在记录，不存在则插入
 * @author: shalom
 * @date: 2020/12/30 16:59
 */
func (m *BtcBlock) Refresh() (err error) {
	m_ := &BtcBlock{
		Height: m.Height,
	}
	m__, err := m_.One()
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return errors.Wrap(err, "m_.One() error")
	}
	if gorm.IsRecordNotFoundError(err) || (err == nil && m__.Id == 0) {
		err := m.Create()
		if err != nil {
			return errors.Wrap(err, "m.Create() error")
		}
		logger.Logger.Info().Uint("block Height",m.Height).Msg("写入一条新 BtcBlock")
	}
	return nil
}
