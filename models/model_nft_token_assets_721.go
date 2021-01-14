package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

var _ = time.Thursday

//NftTokenAssets721
type NftTokenAssets721 struct {
	Id              uint       `gorm:"column:id" form:"id" json:"id" comment:"自增加主键" sql:"bigint(20),PRI"`
	CreatedAt       *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	ContractAddress string     `gorm:"column:contract_address" form:"contract_address" json:"contract_address" comment:"721合约地址" sql:"varchar(64),MUL"`
	TokenId         float64    `gorm:"column:token_id" form:"token_id" json:"token_id" comment:"tokenID" sql:"double(78,0)"`
	OwnerAddress    string     `gorm:"column:owner_address" form:"owner_address" json:"owner_address" comment:"拥有者地址" sql:"varchar(64),MUL"`
	Uri             string     `gorm:"column:uri" form:"uri" json:"uri" comment:"tokenID_uri" sql:"mediumtext"`
}

//TableName
func (m *NftTokenAssets721) TableName() string {
	return "nft_token_assets_721"
}

//One
func (m *NftTokenAssets721) One() (one *NftTokenAssets721, err error) {
	one = &NftTokenAssets721{}
	err = crudOne(m, one)
	return
}

//All
func (m *NftTokenAssets721) All(q *PaginationQuery) (list *[]NftTokenAssets721, total uint, err error) {
	list = &[]NftTokenAssets721{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *NftTokenAssets721) Update() (err error) {
	where := NftTokenAssets721{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *NftTokenAssets721) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *NftTokenAssets721) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *NftTokenAssets721) Has() (b bool,m__ *NftTokenAssets721, err error) {
	m_ := &NftTokenAssets721{
		ContractAddress: m.ContractAddress,
		TokenId: m.TokenId,
	}
	m__, err = m_.One()
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return false,nil, errors.Wrap(err, "m_.One() error")
	}
	return gorm.IsRecordNotFoundError(err) || (err == nil && m__.Id == 0), m__,nil
}


func (m *NftTokenAssets721)RefreshAssets() error{
	return nil
}