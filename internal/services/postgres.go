package services

import (
	"fmt"
	"log"
	"server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB

func DB() *gorm.DB {
	return _db
}

func LoadDatabase() {
	log.Default().Print("LOADING Database")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		Conf.DB.HOST,
		Conf.DB.USER,
		Conf.DB.PASS,
		Conf.DB.NAME,
		Conf.DB.PORT,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	loadMigrations(database)
}

func loadMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
