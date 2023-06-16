package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Get("/:slug", Redirect)
	app.Post("/shorten", Shorten)
}
