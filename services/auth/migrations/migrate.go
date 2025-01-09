package migrations

import (
	"job-posting/services/auth/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	} else {
		log.Println("Database migrated successfully.")
	}
}
