package database

import (
	"github.com/tedjoskb/go-restapi-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	conn := "host=localhost user=postgres password=admin dbname=go_restapi_fiber port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Book{})
	DB = db
}
