package models

type DBInterface interface {
	Refresh() error
	SetBaseURI(string)
	Has() (bool,error)
	GetBlockNum()(blockNum uint,err error)
}

type AssetsDBInterface interface {
	Has() (b bool, err error)
	RefreshAssets() error
}