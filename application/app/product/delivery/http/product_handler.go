package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"go-fiber-clean-architecture/application/app/product/request"
	"go-fiber-clean-architecture/application/domain"
	"strconv"
)

type ProductHandler struct {
	ProductUseCase domain.ProductUseCase
	Validate        *validator.Validate
}

func NewProductHandler(productUseCase domain.ProductUseCase, validate *validator.Validate) domain.ProductHandler {
	return &ProductHandler{
		ProductUseCase: productUseCase,
		Validate:        validate,
	}
}

func (handler ProductHandler) GetByID(c *fiber.Ctx) error {
	// ini adalah contoh penggunaan error handling
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// mengambil data dari usecase
	res, err := handler.ProductUseCase.GetByID(c.Context(), int64(id))

	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success get product",
		"data":    res,
	})
}

func (handler ProductHandler) Create(c *fiber.Ctx) error {
	var product domain.Product
	var productCreateRequest request.ProductCreateRequest
	// ambil data dari request ke struct
	c.BodyParser(&productCreateRequest)
	err := handler.Validate.Struct(productCreateRequest)
	if err != nil {
		return err
	}

	copier.Copy(&product, productCreateRequest)

	//insert data ke database menggunakan usecase
	res, err := handler.ProductUseCase.Create(c.Context(), product)

	if err != nil {
		return err
	}

	// return response
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "Success create product",
		"data":    res,
	})
}

func (handler *ProductHandler) GetAll(c *fiber.Ctx) error {
	res, err := handler.ProductUseCase.GetAll(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"message": "Success get all product",
		"data":    res,
	})
}

func (handler *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// delete data ke database menggunakan usecase
	err = handler.ProductUseCase.Delete(c.Context(), int64(id))
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success delete product",
	})
}

func (handler *ProductHandler) Update(c *fiber.Ctx) error {
	var product domain.Product
	var productUpdateRequest request.ProductUpdateRequest
	// ambil data dari request ke struct
	c.BodyParser(&productUpdateRequest)
	err := handler.Validate.Struct(productUpdateRequest)

	if err != nil {
		return err
	}

	mapstructure.Decode(productUpdateRequest, &product)

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

	product.ID = int64(id)

	// update data ke database menggunakan usecase
	res, err := handler.ProductUseCase.Update(c.Context(), product)
	if err != nil {
		return err
	}

	// return response
	return c.JSON(map[string]any{
		"message": "Success update product",
		"data":    res,
	})
}