package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//BtcBlockTx
type BtcBlockTx struct {
	Id      uint `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20) unsigned,PRI"`
	BlockId uint `gorm:"column:block_id" form:"block_id" json:"block_id" comment:"block id" sql:"bigint(20) unsigned,MUL"`
	TxId    uint `gorm:"column:tx_id" form:"tx_id" json:"tx_id" comment:"transaction id" sql:"bigint(20) unsigned"`
}

//TableName
func (m *BtcBlockTx) TableName() string {
	return "btc_block_tx"
}

//One
func (m *BtcBlockTx) One() (one *BtcBlockTx, err error) {
	one = &BtcBlockTx{}
	err = crudOne(m, one)
	return
}

//All
func (m *BtcBlockTx) All(q *PaginationQuery) (list *[]BtcBlockTx, total uint, err error) {
	list = &[]BtcBlockTx{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *BtcBlockTx) Update() (err error) {
	where := BtcBlockTx{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *BtcBlockTx) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *BtcBlockTx) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *BtcBlockTx) Refresh() (err error) {
	m_ := &BtcBlockTx{
		TxId: m.TxId,
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
		logger.Logger.Info().Uint("tx TxID",m.TxId).Msg("写入一条新 BtcBlockTx")
	}
	return nil
}