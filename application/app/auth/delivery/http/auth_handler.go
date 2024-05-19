package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/domain"
)

type AuthHandler struct {
	Validate *validator.Validate
	Ucase domain.AuthUseCase
}

func NewAuthHandler(authUseCase domain.AuthUseCase, validate *validator.Validate) domain.AuthHandler{
	return &AuthHandler{
		Validate: validate,
		Ucase: authUseCase,
	}
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
	data, err := handler.Ucase.Login(c, authCreateRequest)
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (handler *AuthHandler) User(c *fiber.Ctx) error {
	claims := handler.Ucase.User(c)
	return c.JSON(fiber.Map{
		"message": "Welcome " + claims["username"].(string),
	})
}

func (handler *AuthHandler) Logout(c *fiber.Ctx) error {
	handler.Ucase.Logout(c)
	return c.JSON(map[string]any{
		"message": "Success logout",
	})
}
