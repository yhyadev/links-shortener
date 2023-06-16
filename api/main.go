package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yhyadev/links-shortener/routes"
)

func main() {
  app := fiber.New()

  routes.Setup(app)

  app.Listen(":8080")
}
