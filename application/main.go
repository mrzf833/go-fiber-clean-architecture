package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app"
	"time"
)

func main()  {
	app.RunServer()
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