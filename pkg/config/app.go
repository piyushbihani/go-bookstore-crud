package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB

func ConnectDataBase() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	if dbUser == "" || dbPass == "" {
		panic("Error: Environment variables DBUserName or DBPassword are not set")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
