package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	e := godotenv.Load("local.env")
	if e != nil {
		panic(e)
	}
	URL := os.Getenv("DB_URL")
	d, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
