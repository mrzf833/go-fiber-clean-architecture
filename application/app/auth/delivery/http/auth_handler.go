package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/middleware"
)

type AuthHandler struct {
	Validate *validator.Validate
	ucase domain.AuthUseCase
}

func NewAuthHandler(app fiber.Router, authUseCase domain.AuthUseCase, validate *validator.Validate) domain.AuthHandler{
	handler := &AuthHandler{
		Validate: validate,
		ucase: authUseCase,
	}

	// setup routes
	app.Post("/login", handler.Login)
	app.Get("/user", append(middleware.AuthMiddleware(), handler.User)...)
	app.Post("/logout", append(middleware.AuthMiddleware(), handler.Logout)...)

	return handler
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var authCreateRequest request.AuthCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&authCreateRequest)
	err := handler.Validate.Struct(authCreateRequest)

	if err != nil {
		return err
	}

	// mengambil data dari usecase
	data, err := handler.ucase.Login(c, authCreateRequest)
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (handler *AuthHandler) User(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.SendString("Welcome " + claims["username"].(string))
}

func (handler *AuthHandler) Logout(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	// delete token di redis
	config.RedisDb.Delete(username.(string))

	return c.JSON(map[string]any{
		"message": "Success logout",
	})
}
