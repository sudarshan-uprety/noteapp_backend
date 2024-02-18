package initializers

import (
	"log"
	"noteapp/database"
	"noteapp/models"
)

func SyncDatabase() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to perform auto migration: %v", err)
	}
}
