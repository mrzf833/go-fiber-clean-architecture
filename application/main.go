package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/exception"
	application "go-fiber-clean-architecture/application/router"
	"log"
)

func main()  {
	validate := validator.New()
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
	application.SetupRouters(app, validate)

	// handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	})

	// run server
	log.Fatal(app.Listen(config.AppUrl + ":" + config.AppPort))
}