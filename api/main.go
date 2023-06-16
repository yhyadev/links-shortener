package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/yhyadev/links-shortener/database"
	"github.com/yhyadev/links-shortener/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	database.Connect()
	routes.Setup(app)

	app.Listen(":8080")
}
