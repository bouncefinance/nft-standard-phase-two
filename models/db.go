package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
	"nft_standard/config"
	"time"
)

var mysqlDB *gorm.DB

func init() {
	//init mysql
	var (
		dbName, user, password, host string
	)
	sec, err := config.Cfg.GetSection(config.DBSection)
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	charset := sec.Key("CHARSET").String()
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, dbName, charset)
	if db, err := gorm.Open("mysql", conn); err == nil {
		mysqlDB = db
	} else {
		logrus.WithError(err).Fatalln("initialize mysql database failed")
	}

	mysqlDB.DB().SetMaxIdleConns(10)
	mysqlDB.DB().SetMaxOpenConns(100)
	mysqlDB.DB().SetConnMaxLifetime(time.Minute * 5)
}

func Close() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
