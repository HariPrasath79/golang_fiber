package database

import (
	"fiber/golang_fiber/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DbConnect() {

	dbURl, has := os.LookupEnv("DB_URL")
	if !has {
		panic("DB_key not found")
	}

	db, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed")
	}
	log.Print("Connection succesful, ok")

	db.AutoMigrate(new(model.Blog))

	DB = db

}
