package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App, validate *validator.Validate)  {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	categoryRouterApi(api, validate)
	authRouterApi(api, validate)
}
