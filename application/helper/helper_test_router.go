package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/exception"
)

type HelperRouter struct {
	Handlers []fiber.Handler
	Method string
	Path string
}

func TestApp(router ...HelperRouter) *fiber.App {
	//validate := validator.New()
	// setup application in main.go
	app := fiber.New(fiber.Config{
		//Prefork: true,

		// ini adalah custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return exception.HandleError(c, err)
		},
	})
	// setup middleware
	app.Use(cors.New(), recover.New())

	// connect to database
	config.ConnectDB()

	// setup routes
	for _, r := range router {
		app.Add(r.Method, r.Path, r.Handlers...)
	}

	// handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	})

	return app
}