package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DSN
const DSN = "camp:123456@(139.196.222.16:3306)/camp?charset=utf8&parseTime=True&loc=Local"
//指定驱动
const DRIVER = "mysql"
//	db, err := gorm.Open("mysql", "ele:123456@(139.196.222.16:3306)/ele?charset=utf8&parseTime=True&loc=Local")

var db *gorm.DB

func init() {
	var err error
	db,err = gorm.Open(DRIVER,DSN)
	if err != nil{
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//defer db.Close()
}
//func Close()  {
//	defer db.Close()
//}

