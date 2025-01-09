package migrations

import (
	"job-posting/services/company/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Company{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	} else {
		log.Println("Database migrated successfully.")
	}
}
