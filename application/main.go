package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/exception"
	application "go-fiber-clean-architecture/application/router"
	"log"
	"time"
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


//func main(){
//	app := fiber.New()
//
//	// Login route
//	app.Post("/api/login", login)
//
//	// Unauthenticated route
//	app.Get("/api", accessible)
//
//	// JWT Middleware
//	app.Use(jwtware.New(jwtware.Config{
//		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
//	}))
//
//	// Restricted Routes
//	app.Get("/api/restricted", middleware.Restricted)
//
//	app.Listen(":3000")
//}

// mencoba middleware secafra langsung
func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}