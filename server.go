package main

import (
	"fiber/golang_fiber/database"
	"fiber/golang_fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DbConnect()
}

func main() {

	db, err := database.DB.DB()

	if err != nil {
		panic("Database didn't connected")
	}

	defer db.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":8080")
}
