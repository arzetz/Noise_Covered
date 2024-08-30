package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "arzetz:8499k8499k@tcp(127.0.0.1:3306)/noize_covered?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB {
	return db
}
