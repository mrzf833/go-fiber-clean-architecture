package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/exception"
	"go-fiber-clean-architecture/application/helper"
	application "go-fiber-clean-architecture/application/router"
)

func AppInit() *fiber.App {
	// validate
	validate := validator.New(validator.WithRequiredStructEnabled())
	// setup application in main.go
	app := fiber.New(fiber.Config{
		//Prefork: true,

		// ini adalah custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return exception.HandleError(c, err)
		},
	})
	// add custom validation
	app.Use(func(ctx *fiber.Ctx) error {
		customValidation := helper.NewCustomValidation(validate, ctx)
		customValidation.RegisterCustomValidation()
		return ctx.Next()
	})
	// setup middleware
	app.Use(cors.New(), recover.New())

	// connect to database
	helper.ConnectDB()
	// connect to redis
	helper.ConnectRedis()

	// setup static file
	app.Static("/static", helper.GetApplicationPath() + "/storage/public")
	// setup routes
	application.SetupRouters(app, validate)

	// handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	})
	return app
}

func RunServer() {
	app := AppInit()
	// run server
	log.Fatal(app.Listen(config.AppUrl + ":" + config.AppPort))
}