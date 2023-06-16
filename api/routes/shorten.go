package routes

import (
	"math/rand"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/yhyadev/links-shortener/database"
	"github.com/yhyadev/links-shortener/helpers"
)

type requestBody struct {
	URL string `json:"url"`
}

func Shorten(ctx *fiber.Ctx) error {
	body := requestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	if govalidator.IsURL(body.URL) != true {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	slug := helpers.Encode(rand.Uint64())
	url := helpers.EnforceHTTP(body.URL)

	link := &database.Link{
		Slug:     slug,
		Redirect: url,
	}

	if err := mgm.Coll(link).Create(link); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(link)
}
