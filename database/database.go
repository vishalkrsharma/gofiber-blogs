package database

import (
	"log"

	"github.com/vishalkrsharma/go-blogs/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=vishal password=1234 dbname=gofiber port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("DB connection failed")
	}

	log.Println("DB connected")

	db.AutoMigrate(new(model.Blog))

	DBConn = db
}
