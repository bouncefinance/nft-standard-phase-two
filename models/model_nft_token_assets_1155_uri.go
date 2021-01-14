package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//NftTokenAssets1155Uri
type NftTokenAssets1155Uri struct {
	Id              uint       `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20),PRI"`
	CreatedAt       *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	ContractAddress string     `gorm:"column:contract_address" form:"contract_address" json:"contract_address" comment:"1155合约地址" sql:"varchar(64),MUL"`
	TokenId         float64    `gorm:"column:token_id" form:"token_id" json:"token_id" comment:"tokenID" sql:"double(78,0)"`
	Uri             string     `gorm:"column:uri" form:"uri" json:"uri" comment:"tokenID_uri" sql:"mediumtext"`
}

//TableName
func (m *NftTokenAssets1155Uri) TableName() string {
	return "nft_token_assets_1155_uri"
}

//One
func (m *NftTokenAssets1155Uri) One() (one *NftTokenAssets1155Uri, err error) {
	one = &NftTokenAssets1155Uri{}
	err = crudOne(m, one)
	return
}

//All
func (m *NftTokenAssets1155Uri) All(q *PaginationQuery) (list *[]NftTokenAssets1155Uri, total uint, err error) {
	list = &[]NftTokenAssets1155Uri{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *NftTokenAssets1155Uri) Update() (err error) {
	where := NftTokenAssets1155Uri{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *NftTokenAssets1155Uri) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *NftTokenAssets1155Uri) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
