package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {

	var err error
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/skr")
	if err != nil {
		log.Panicln("mysql connect err:", err.Error())
	}

}
