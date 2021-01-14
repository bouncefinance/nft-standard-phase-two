package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//BtcTx
type BtcTx struct {
	Id            uint   `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20) unsigned,PRI"`
	Version       uint   `gorm:"column:version" form:"version" json:"version" comment:"版本" sql:"int(10) unsigned"`
	VSize         uint   `gorm:"column:v_size" form:"v_size" json:"v_size" comment:"虚拟交易大小" sql:"int(10) unsigned"`
	LockTime      uint   `gorm:"column:lock_time" form:"lock_time" json:"lock_time" comment:"lock time" sql:"int(10) unsigned"`
	BlockHash     string `gorm:"column:block_hash" form:"block_hash" json:"block_hash" comment:"区块Hash" sql:"varchar(64)"`
	TxId          string `gorm:"column:tx_id" form:"tx_id" json:"tx_id" comment:"交易ID，存入block的就是这个" sql:"varchar(64),UNI"`
	Hash          string `gorm:"column:hash" form:"hash" json:"hash" comment:"交易Hash" sql:"varchar(64)"`
	Size          uint   `gorm:"column:size" form:"size" json:"size" comment:"序列化后的交易数据大小" sql:"int(10) unsigned"`
	Weight        uint   `gorm:"column:weight" form:"weight" json:"weight" comment:"交易权重" sql:"int(10) unsigned"`
	Confirmations uint   `gorm:"column:confirmations" form:"confirmations" json:"confirmations" comment:"交易确认数" sql:"int(10) unsigned"`
	Time          uint   `gorm:"column:time" form:"time" json:"time" comment:"时间戳，和block的value一样" sql:"int(10) unsigned"`
	BlockTime     uint   `gorm:"column:block_time" form:"block_time" json:"block_time" comment:"区块时间戳" sql:"int(10) unsigned"`
}

//TableName
func (m *BtcTx) TableName() string {
	return "btc_tx"
}

//One
func (m *BtcTx) One() (one *BtcTx, err error) {
	one = &BtcTx{}
	err = crudOne(m, one)
	return
}

//All
func (m *BtcTx) All(q *PaginationQuery) (list *[]BtcTx, total uint, err error) {
	list = &[]BtcTx{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *BtcTx) Update() (err error) {
	where := BtcTx{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *BtcTx) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *BtcTx) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *BtcTx) Refresh() (err error) {
	m_ := &BtcTx{
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
		logger.Logger.Info().Str("tx TxID",m.TxId).Msg("写入一条新 BtcTx")
	}
	return nil
}