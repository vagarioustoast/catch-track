package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		log.Fatal("DB_PASSWORD environment variable is required but not set")
	}

	dsn := fmt.Sprintf("host=localhost user=gorm password=%s dbname=gorm port=9920 sslmode=disable TimeZone=America/Los_Angeles", password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}
