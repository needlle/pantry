package server

import (
	"pantry/api"

	"github.com/gofiber/fiber/v2"
)

type App interface {
	Listen(string) error
}

func New() App {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy!")
	})

	v1 := app.Group("/api/v1")
	api.Setup(v1)

	return app
}
