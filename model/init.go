package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"skrshop-api/conf"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(conf.DBType, conf.DBSource)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
