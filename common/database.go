package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := 3306
	datebase := "test"
	username := "root"
	password := "12345678"
	charset := "utf8mb4"
	//args:=fmt.Sprintf('%s:%s@tcp(%s:%s)/%s?charset=%s')
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, datebase, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("发现错误" + err.Error())
	}
	return db
}

func GetDB() *gorm.DB {
	return DB
}
