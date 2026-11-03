package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "mysql://21BxSUYmz82NfuY.root:LnpFsVWimAL6Uwim@gateway01.eu-central-1.prod.aws.tidbcloud.com:4000/test?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
