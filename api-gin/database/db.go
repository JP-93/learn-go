package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-gin/models"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	connection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatal("Error connect database", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
