package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"m/v2/models"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Error connect database", err)
	}
	DB.AutoMigrate(&models.Lists)
}
