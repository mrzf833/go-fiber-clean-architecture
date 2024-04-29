package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"go-fiber-clean-architecture/application/app/file_save/request"
	"go-fiber-clean-architecture/application/domain"
)

type FileSaveHandler struct {
	FileSaveUseCase domain.FileSaveUsecase
	Validate        *validator.Validate
}

func NewFileSaveHandler(app fiber.Router, fileSaveUseCase domain.FileSaveUsecase, validate *validator.Validate) {
	handler := &FileSaveHandler{
		FileSaveUseCase: fileSaveUseCase,
		Validate:        validate,
	}

	// setup routes
	//app.Get("/:id", handler.GetByID)
	//app.Get("/", handler.GetAll)
	app.Post("/", handler.Create)
}

func (handler *FileSaveHandler) Create(c *fiber.Ctx) error {
	var fileSave domain.FileSave
	var fileSaveCreateRequest request.FileSaveCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&fileSaveCreateRequest)
	err := handler.Validate.Struct(fileSaveCreateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(fileSaveCreateRequest, &fileSave)

	//insert data ke database menggunakan usecase
	//res, err := handler.FileSaveUseCase.Create(c, fileSave)

	//if err != nil {
	//	return err
	//}

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create file save",
		//"data":    res,
	})
}