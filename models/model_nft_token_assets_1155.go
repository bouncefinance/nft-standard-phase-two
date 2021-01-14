package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//NftTokenAssets1155
type NftTokenAssets1155 struct {
	Id              uint       `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20),PRI"`
	CreatedAt       *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	ContractAddress string     `gorm:"column:contract_address" form:"contract_address" json:"contract_address" comment:"1155合约地址" sql:"varchar(64),MUL"`
	TokenId         float64    `gorm:"column:token_id" form:"token_id" json:"token_id" comment:"tokenID" sql:"double(78,0)"`
	OwnerAddress    string     `gorm:"column:owner_address" form:"owner_address" json:"owner_address" comment:"拥有者地址" sql:"varchar(64),MUL"`
	Balance         float64    `gorm:"column:balance" form:"balance" json:"balance" comment:"拥有余额" sql:"double(78,0)"`
	Uri             string     `gorm:"column:uri" form:"uri" json:"uri" comment:"tokenID_uri" sql:"mediumtext"`
}

//TableName
func (m *NftTokenAssets1155) TableName() string {
	return "nft_token_assets_1155"
}

//One
func (m *NftTokenAssets1155) One() (one *NftTokenAssets1155, err error) {
	one = &NftTokenAssets1155{}
	err = crudOne(m, one)
	return
}

//All
func (m *NftTokenAssets1155) All(q *PaginationQuery) (list *[]NftTokenAssets1155, total uint, err error) {
	list = &[]NftTokenAssets1155{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *NftTokenAssets1155) Update() (err error) {
	where := NftTokenAssets1155{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *NftTokenAssets1155) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *NftTokenAssets1155) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
