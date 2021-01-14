package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"nft_standard/logger"
	"time"
)

var _ = time.Thursday

//NftContractRecord721
type NftContractRecord721 struct {
	Id              uint       `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20),PRI"`
	CreatedAt       *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	ContractAddress string     `gorm:"column:contract_address" form:"contract_address" json:"contract_address" comment:"721合约地址" sql:"varchar(64),UNI"`
	BlockNumber     uint       `gorm:"column:block_number" form:"block_number" json:"block_number" comment:"合约第一次出现块高" sql:"int(10) unsigned,MUL"`
	BaseUri         string     `gorm:"column:base_uri" form:"base_uri" json:"base_uri" comment:"base_uri" sql:"varchar(240)"`
}

//TableName
func (m *NftContractRecord721) TableName() string {
	return "nft_contract_record_721"
}

//One
func (m *NftContractRecord721) One() (one *NftContractRecord721, err error) {
	one = &NftContractRecord721{}
	err = crudOne(m, one)
	return
}



//All
func (m *NftContractRecord721) All(q *PaginationQuery) (list *[]NftContractRecord721, total uint, err error) {
	list = &[]NftContractRecord721{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *NftContractRecord721) Update() (err error) {
	where := NftContractRecord721{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *NftContractRecord721) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *NftContractRecord721) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

/**
 * @parameter:
 * @return:
 * @Description: 判断数据库中是否存在记录，不存在则插入，存在则更新
 * @author: shalom
 * @date: 2020/12/30 16:59
 */
func (m *NftContractRecord721) Refresh() (err error) {
	m_ := &NftContractRecord721{
		ContractAddress: m.ContractAddress,
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
		logger.Logger.Info().Str("contract address", m.ContractAddress).Msg("写入一条新721 contract record")
	} else if m__.BlockNumber != m.BlockNumber || m__.BaseUri != m.BaseUri {
		err := m.Update()
		if err != nil {
			return errors.Wrap(err, "m.Update() error")
		}
		logger.Logger.Info().Str("contract address", m.ContractAddress).Msg("更新一条新721 contract record")
	}
	return nil
}

func (m *NftContractRecord721) SetBaseURI(uri string) {
	m.BaseUri = uri
}
func (m *NftContractRecord721) Has() (b bool, err error) {
	m_ := &NftContractRecord721{
		ContractAddress: m.ContractAddress,
	}
	m__, err := m_.One()
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return false, errors.Wrap(err, "m_.One() error")
	}
	return gorm.IsRecordNotFoundError(err) || (err == nil && m__.Id == 0), nil
}
//	利用覆盖索引，查询指定字段blockNum
func (m *NftContractRecord721) GetBlockNum()(blockNum uint,err error) {
	one := &NftContractRecord721{}
	err = getSelected([]string{"block_number"},m,one)
	blockNum = one.BlockNumber
	return
}
