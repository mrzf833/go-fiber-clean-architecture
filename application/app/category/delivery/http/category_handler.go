package http

import (
	"fmt"
	"go-fiber-clean-architecture/application/app/category/request"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type CategoryHandler struct {
	CategoryUseCase domain.CategoryUseCase
	Validate        *validator.Validate
}

func NewCategoryHandler(categoryUseCase domain.CategoryUseCase, validate *validator.Validate) domain.CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: categoryUseCase,
		Validate:        validate,
	}
}

func (handler *CategoryHandler) GetByID(c *fiber.Ctx) error {

	// ini adalah contoh penggunaan error handling
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// mengambil data dari usecase
	res, err := handler.CategoryUseCase.GetByID(c.Context(), int64(id))

	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success get category",
		"data":    res,
	})
}

func (handler *CategoryHandler) GetAll(c *fiber.Ctx) error {
	res, err := handler.CategoryUseCase.GetAll(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"message": "Success get all category",
		"data":    res,
	})
}

func (handler *CategoryHandler) Create(c *fiber.Ctx) error {
	var category domain.Category
	var categoryCreateRequest request.CategoryCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&categoryCreateRequest)
	err := handler.Validate.Struct(categoryCreateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(categoryCreateRequest, &category)

	//insert data ke database menggunakan usecase
	res, err := handler.CategoryUseCase.Create(c.Context(), category)

	if err != nil {
		return err
	}

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create category",
		"data":    res,
	})
}

func (handler *CategoryHandler) Update(c *fiber.Ctx) error {
	var category domain.Category
	var categoryCreateRequest request.CategoryCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&categoryCreateRequest)
	err := handler.Validate.Struct(categoryCreateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(categoryCreateRequest, &category)

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

	category.ID = int64(id)

	// update data ke database menggunakan usecase
	res, err := handler.CategoryUseCase.Update(c.Context(), category)
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success update category",
		"data":    res,
	})
}

func (handler *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// delete data ke database menggunakan usecase
	err = handler.CategoryUseCase.Delete(c.Context(), int64(id))
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success delete category",
	})
}

func (handler *CategoryHandler) CreateWithCsv(c *fiber.Ctx) error {
	var categoryCreateCsvRequest request.CategoryCreateCsvRequest
	// ambil data dari request ke struct
	categoryCreateCsvRequest.File,_ = c.FormFile("file")
	c.BodyParser(&categoryCreateCsvRequest)
	err := handler.Validate.Struct(categoryCreateCsvRequest)
	if err != nil {
		return err
	}

	openFile, err := categoryCreateCsvRequest.File.Open()
	if err != nil {
		return err
	}

	// ini untuk ngetracking berapa lama proses insert data ke database menggunakan go routine dan gorm per batches
	trackerCategory := domain.TrackerCategory{
		Name: "Category Berubah",
		Now: strconv.FormatInt(time.Now().UnixMilli(), 10),
		End: " ",
	}
	config.DB.Create(&trackerCategory)

	//insert data ke database menggunakan usecase
	go handler.CategoryUseCase.CreateWithCsv(c.Context(), openFile, trackerCategory.ID)

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create category with csv",
	})
}

func (handler *CategoryHandler) CreateWithCsvQueue(c *fiber.Ctx) error {
	var categoryCreateCsvRequest request.CategoryCreateCsvRequest
	// ambil data dari request ke struct
	categoryCreateCsvRequest.File,_ = c.FormFile("file")
	c.BodyParser(&categoryCreateCsvRequest)
	err := handler.Validate.Struct(categoryCreateCsvRequest)
	if err != nil {
		return err
	}

	openFile, err := categoryCreateCsvRequest.File.Open()
	if err != nil {
		return err
	}

	//insert data ke database menggunakan usecase
	err = handler.CategoryUseCase.CreateWithCsvQueue(c.Context(), openFile)

	if err != nil{
		fmt.Println("disini")
		return err
	}

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create category with csv on queue",
	})
}
