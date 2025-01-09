package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init databse tapi konfigurasi menggukanan env
func DBInit() *gorm.DB {

	// dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	

	dsn := "host="+dbHost+" user="+dbUserName+" password="+dbPassword+" dbname="+dbName+" port="+dbPort+" sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}