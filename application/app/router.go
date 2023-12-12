package app

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/category/delivery/http"
	"go-fiber-clean-architecture/application/app/category/usecase"
)

func SetupRouters(app *fiber.App)  {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// setup category router
	http.NewCategoryHandler(api, usecase.NewCategoryUseCase())
}
