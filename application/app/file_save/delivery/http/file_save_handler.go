package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"go-fiber-clean-architecture/application/app/file_save/request"
	"go-fiber-clean-architecture/application/domain"
	"strconv"
)

type FileSaveHandler struct {
	FileSaveUseCase domain.FileSaveUsecase
	Validate        *validator.Validate
}

func NewFileSaveHandler(app fiber.Router, fileSaveUseCase domain.FileSaveUsecase, validate *validator.Validate) domain.FileSaveHandler {
	handler := &FileSaveHandler{
		FileSaveUseCase: fileSaveUseCase,
		Validate:        validate,
	}

	// setup routes
	app.Get("/:id", handler.GetByID)
	app.Get("/", handler.GetAll)
	app.Post("/", handler.Create)
	app.Post("/:id", handler.Update)
	app.Delete("/:id", handler.Delete)

	return handler
}

func (handler *FileSaveHandler) GetAll(c *fiber.Ctx) error {
	res, err := handler.FileSaveUseCase.GetAll(c)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"message": "Success get all file save",
		"data":    res,
	})
}

func (handler *FileSaveHandler) GetByID(c *fiber.Ctx) error {

	// ini adalah contoh penggunaan error handling
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// mengambil data dari usecase
	res, err := handler.FileSaveUseCase.GetByID(c, int64(id))

	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success get file save",
		"data":    res,
	})
}

func (handler *FileSaveHandler) Create(c *fiber.Ctx) error {
	var fileSave domain.FileSave
	var fileSaveCreateRequest request.FileSaveCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&fileSaveCreateRequest)
	fileSaveCreateRequest.File, _ = c.FormFile("file")
	err := handler.Validate.Struct(fileSaveCreateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(fileSaveCreateRequest, &fileSave)

	//insert data ke database menggunakan usecase
	res, err := handler.FileSaveUseCase.Create(c, fileSave)

	if err != nil {
		return err
	}

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create file save",
		"data":    res,
	})
}

func (handler *FileSaveHandler) Update(c *fiber.Ctx) error {
	var fileSave domain.FileSave
	var fileSaveCreateRequest request.FileSaveCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&fileSaveCreateRequest)
	fileSaveCreateRequest.File, _ = c.FormFile("file")
	err := handler.Validate.Struct(fileSaveCreateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(fileSaveCreateRequest, &fileSave)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	fileSave.ID = int64(id)

	// update data ke database menggunakan usecase
	res, err := handler.FileSaveUseCase.Update(c, fileSave)
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success update file save",
		"data":    res,
	})
}

func (handler *FileSaveHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// delete data ke database menggunakan usecase
	err = handler.FileSaveUseCase.Delete(c, int64(id))
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success delete file save",
	})
}
