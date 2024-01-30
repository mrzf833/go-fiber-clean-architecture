package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/middleware"
	"time"
)

type AuthHandler struct {
	Validate *validator.Validate
	ucase domain.AuthUseCase
}

func NewAuthHandler(app fiber.Router, authUseCase domain.AuthUseCase, validate *validator.Validate) {
	handler := &AuthHandler{
		Validate: validate,
		ucase: authUseCase,
	}

	// setup routes
	app.Post("/login", handler.Login)
	app.Get("/restricted", middleware.AuthMiddleware(),  handler.Restricted)
	app.Post("/logout", middleware.AuthMiddleware(),  handler.Logout)
}

//func login(c *fiber.Ctx) error {
//
//	jwt.MapClaims{}
//
//	jwt.NewWithClaims(jwt.SigningMethodHS256)
//	return nil
//}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var authCreateRequest request.AuthCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&authCreateRequest)
	err := handler.Validate.Struct(authCreateRequest)

	if err != nil {
		return err
	}

	// mengambil data dari usecase
	data, err := handler.ucase.Login(c.Context(), authCreateRequest)
	if err != nil {
		return err
	}

	return c.JSON(data)
}


func (handler *AuthHandler) Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	//name := claims["username"]
	//return c.SendString("Welcome " + name)
	return c.JSON(claims)
}

func (handler *AuthHandler) Logout(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"]
	expired := time.Now().Add(-config.ExpireToken)
	claims["exp"] =expired.Unix()

	c.Cookie(&fiber.Cookie{
		Name:     "user",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 24)),
	})

	return c.JSON(map[string]any{
		"message": "Success logout",
		"date": map[string]any{
			"username": username,
			"exp": claims["exp"],
		},
	})

	//expired := time.Now().Add(-config.ExpireToken)
	//c.Cookie(&fiber.Cookie{
	//	Name:    "token",
	//	Value:   "",
	//	Expires: expired,
	//})
	//
	//return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
