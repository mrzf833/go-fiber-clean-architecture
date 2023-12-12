package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	app2 "go-fiber-clean-architecture/application/app"
	"go-fiber-clean-architecture/application/config"
	"log"
)

func main()  {
	// setup application in main.go
	app := fiber.New()
	// setup middleware
	app.Use(cors.New(), recover.New())

	app2.SetupRouters(app)

	// handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	})

	log.Fatal(app.Listen(config.Config("APP_URL") + ":" + config.Config("APP_PORT")))
}