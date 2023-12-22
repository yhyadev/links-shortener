package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/yhyadev/links-shortener/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Redirect(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	link := &database.Link{}
	if err := mgm.Coll(link).First(bson.M{"slug": slug}, link); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid slug"})
	}

	return ctx.Redirect(link.Redirect, 301)
}
