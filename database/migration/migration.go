package migration

import (
	"fmt"
	"log"

	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/models"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&models.Users{}, &models.Book{}, &models.Account{}, &models.Entries{}, &models.Transfer{})

	if err == nil {
		fmt.Println("Database Migrated")
	} else if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Migration Failed")
	}

}
