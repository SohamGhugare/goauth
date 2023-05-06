package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database...")
	}
	DB = db

}
