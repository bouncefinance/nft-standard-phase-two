package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

var _ = time.Thursday

//NftContractRecord
type NftContractRecord struct {
	Id              uint   `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20),PRI"`
	ContractAddress string `gorm:"column:contract_address" form:"contract_address" json:"contract_address" comment:"合约地址" sql:"varchar(64),MUL"`
}

//TableName
func (m *NftContractRecord) TableName() string {
	return "nft_contract_record"
}

//One
func (m *NftContractRecord) One() (one *NftContractRecord, err error) {
	one = &NftContractRecord{}
	err = crudOne(m, one)
	return
}

//All
func (m *NftContractRecord) All(q *PaginationQuery) (list *[]NftContractRecord, total uint, err error) {
	list = &[]NftContractRecord{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *NftContractRecord) Update() (err error) {
	where := NftContractRecord{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *NftContractRecord) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *NftContractRecord) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *NftContractRecord) Has() (b bool, err error) {
	m__, err := m.One()
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return false, errors.Wrap(err, "m.One() error")
	}
	return !(gorm.IsRecordNotFoundError(err) || (err == nil && m__.Id == 0)), nil
}


