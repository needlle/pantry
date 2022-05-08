package server

import (
	"pantry/api"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy!")
	})

	v1 := app.Group("/api/v1")
	api.Setup(v1)

	return app
}
