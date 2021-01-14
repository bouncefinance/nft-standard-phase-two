package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//EthTx
type EthTx struct {
	Id          uint    `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20) unsigned,PRI"`
	BlockNumber uint    `gorm:"column:block_number" form:"block_number" json:"block_number" comment:"所在块高" sql:"int(10) unsigned"`
	BlockHash   string  `gorm:"column:block_hash" form:"block_hash" json:"block_hash" comment:"区块Hash" sql:"varchar(64)"`
	From        string  `gorm:"column:from" form:"from" json:"from" comment:"from" sql:"varchar(64)"`
	To          string  `gorm:"column:to" form:"to" json:"to" comment:"to" sql:"varchar(64)"`
	Gas         uint    `gorm:"column:gas" form:"gas" json:"gas" comment:"消耗gas" sql:"int(10) unsigned"`
	GasPrice    float64 `gorm:"column:gas_price" form:"gas_price" json:"gas_price" comment:"gasPrice" sql:"double(78,0)"`
	Hash        string  `gorm:"column:hash" form:"hash" json:"hash" comment:"交易Hash" sql:"varchar(64),UNI"`
	Nonce       uint    `gorm:"column:nonce" form:"nonce" json:"nonce" comment:"nonce" sql:"int(10) unsigned"`
	Value       float64 `gorm:"column:value" form:"value" json:"value" comment:"交易value" sql:"double(78,0)"`
}

//TableName
func (m *EthTx) TableName() string {
	return "eth_tx"
}

//One
func (m *EthTx) One() (one *EthTx, err error) {
	one = &EthTx{}
	err = crudOne(m, one)
	return
}

//All
func (m *EthTx) All(q *PaginationQuery) (list *[]EthTx, total uint, err error) {
	list = &[]EthTx{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *EthTx) Update() (err error) {
	where := EthTx{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *EthTx) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *EthTx) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *EthTx) Refresh() (err error) {
	m_ := &EthTx{
		Hash: m.Hash,
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
		logger.Logger.Info().Interface("transaction",m).Msg("写入一条新 EthTx")
	}
	return nil
}