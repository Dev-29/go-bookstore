package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "Devabhiram@2003"
const DB_NAME = "go_db"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var db *gorm.DB

func Connect() {
	d, err := gorm.Open(mysql.Open(DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
