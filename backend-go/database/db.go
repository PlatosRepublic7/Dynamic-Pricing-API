package database

import (
	"dynamic-pricing-api/backend-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=pricing port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	log.Println("Database connection established")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Database migration failed: ", err)
	}
	log.Println("Database migration completed")
}
