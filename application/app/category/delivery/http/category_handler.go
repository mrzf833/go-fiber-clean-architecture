package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/domain"
)

type CategoryHandler struct {
	CategoryUseCase domain.CategoryUseCase
}

func NewCategoryHandler(app fiber.Router, categoryUseCase domain.CategoryUseCase) {
	handler := &CategoryHandler{
		CategoryUseCase: categoryUseCase,
	}

	app.Get("/category/:id", handler.GetByID)
}

func (handler *CategoryHandler) GetByID(c *fiber.Ctx) error {
	return c.SendString("Hello, category 👋!")
}