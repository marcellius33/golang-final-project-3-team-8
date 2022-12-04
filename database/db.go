package database

import (
	"log"
	"os"

	"kanbanboard/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_CONFIG")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Task{})
	log.Printf("Success connecting to database")
}

func GetDB() *gorm.DB {
	return db
}
