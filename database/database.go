package database

import (
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
}
