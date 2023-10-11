package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	conn := "host=localhost user=root password=root dbname=simpleBank port=8001 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	// db.AutoMigrate(&models.Book{})
	DB = db
}
