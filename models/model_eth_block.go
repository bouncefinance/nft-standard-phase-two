package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//EthBlock
type EthBlock struct {
	Id         uint    `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20) unsigned,PRI"`
	Hash       string  `gorm:"column:hash" form:"hash" json:"hash" comment:"区块Hash" sql:"varchar(64),UNI"`
	Number     uint    `gorm:"column:number" form:"number" json:"number" comment:"区块高度" sql:"int(10) unsigned,UNI"`
	Difficulty float64 `gorm:"column:difficulty" form:"difficulty" json:"difficulty" comment:"区块难度" sql:"double(78,0)"`
	GasLimit   uint    `gorm:"column:gas_limit" form:"gas_limit" json:"gas_limit" comment:"区块 gas_limit" sql:"int(10) unsigned"`
	GasUsed    uint    `gorm:"column:gas_used" form:"gas_used" json:"gas_used" comment:"区块 gas_used" sql:"int(10) unsigned"`
	Nonce      uint    `gorm:"column:nonce" form:"nonce" json:"nonce" comment:"nonce值" sql:"int(10) unsigned"`
	Size       uint    `gorm:"column:size" form:"size" json:"size" comment:"区块字节数" sql:"int(10) unsigned"`
	Miner      string  `gorm:"column:miner" form:"miner" json:"miner" comment:"miner" sql:"varchar(64)"`
	ParentHash string  `gorm:"column:parent_hash" form:"parent_hash" json:"parent_hash" comment:"前一区块的哈希" sql:"varchar(64)"`
	Time       uint    `gorm:"column:time" form:"time" json:"time" comment:"区块创建时间戳" sql:"int(10) unsigned"`
	NTx        uint    `gorm:"column:n_tx" form:"n_tx" json:"n_tx" comment:"block包含的交易数量" sql:"int(10) unsigned"`
	StateRoot  string  `gorm:"column:state_root" form:"state_root" json:"state_root" comment:"stateRoot" sql:"varchar(64)"`
}

//TableName
func (m *EthBlock) TableName() string {
	return "eth_block"
}

//One
func (m *EthBlock) One() (one *EthBlock, err error) {
	one = &EthBlock{}
	err = crudOne(m, one)
	return
}

//All
func (m *EthBlock) All(q *PaginationQuery) (list *[]EthBlock, total uint, err error) {
	list = &[]EthBlock{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *EthBlock) Update() (err error) {
	where := EthBlock{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *EthBlock) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *EthBlock) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *EthBlock) Refresh() (err error) {
	m_ := &EthBlock{
		Number: m.Number,
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
		logger.Logger.Info().Interface("block", m).Msg("写入一条新 EthBlock")
	}
	return nil
}
