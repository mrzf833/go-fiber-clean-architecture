package http_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpDelivery "go-fiber-clean-architecture/application/app/category/delivery/http"
	"go-fiber-clean-architecture/application/app/category/mocks"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"go-fiber-clean-architecture/application/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllCategoryWithMock(t *testing.T) {
	mockCategoryUseCase := new(mocks.CategoryUsecase)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}
	// buat list mock data
	mockListCategory := make([]domain.Category, 0)
	mockListCategory = append(mockListCategory, mockCategory)

	t.Run("success", func(t *testing.T) {
		mockCategoryUseCase.On("GetAll", mock.Anything).Return(mockListCategory, nil).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
		}


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/category", nil)


		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetAll},
			Method: fiber.MethodGet,
			Path: "/category",
		})


		// panggil handler
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})
}

func TestGetByIdCategoryWithMock(t *testing.T) {
	mockCategoryUseCase := new(mocks.CategoryUsecase)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryUseCase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetByID},
			Method: fiber.MethodGet,
			Path: "/category/:id",
		})


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/category/1", nil)


		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		mockCategoryUseCase.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockCategoryUseCase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Category{}, exception.ErrNotFound).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
		}
		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetByID},
			Method: fiber.MethodGet,
			Path: "/category/:id",
		})


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/category/1", nil)


		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})
}

func TestCreateCategoryWithMock(t *testing.T) {
	mockCategoryUseCase := new(mocks.CategoryUsecase)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	validate := validator.New()

	t.Run("success", func(t *testing.T) {
		mockCategoryUseCase.On("Create", mock.Anything, mock.AnythingOfType("domain.Category")).Return(mockCategory, nil).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
			Validate: validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Create},
			Method: fiber.MethodPost,
			Path: "/category",
		})

		data := strings.NewReader(`{"Name": "Test"}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/category", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})

	t.Run("bad-request", func(t *testing.T) {
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
			Validate: validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Create},
			Method: fiber.MethodPost,
			Path: "/category",
		})

		data := strings.NewReader(`{"Name": ""}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/category", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})
}

func TestUpdateCategoryWithMock(t *testing.T) {
	mockCategoryUseCase := new(mocks.CategoryUsecase)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	validate := validator.New()

	t.Run("success", func(t *testing.T) {
		mockCategoryUseCase.On("Update", mock.Anything, mock.AnythingOfType("domain.Category")).Return(mockCategory, nil).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
			Validate: validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Update},
			Method: fiber.MethodPut,
			Path: "/category/:id",
		})

		// request body
		data := strings.NewReader(`{"ID": 1, "Name": "Test"}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPut, "/category/1", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})

	t.Run("bad-request", func(t *testing.T) {
		mockCategoryUseCase.On("Update", mock.Anything, mock.AnythingOfType("domain.Category")).Return(nil, assert.AnError).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
			Validate:        validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Update},
			Method:   fiber.MethodPut,
			Path:     "/category/:id",
		})

		data := strings.NewReader(`{"ID": 1, "Name": ""}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPut, "/category/1", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)

		//bytes, err := io.ReadAll(res.Body)
		//fmt.Println(string(bytes))
		//fmt.Println("asdasd")
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestDeleteCategoryWithMock(t *testing.T) {
	mockCategoryUseCase := new(mocks.CategoryUsecase)

	t.Run("success", func(t *testing.T) {
		mockCategoryUseCase.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Delete},
			Method: fiber.MethodDelete,
			Path: "/category/:id",
		})

		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/category/1", nil)

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockCategoryUseCase.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(exception.ErrNotFound).Once()
		// buat handler
		handler := httpDelivery.CategoryHandler{
			CategoryUseCase: mockCategoryUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Delete},
			Method: fiber.MethodDelete,
			Path: "/category/:id",
		})

		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/category/1", nil)

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockCategoryUseCase.AssertExpectations(t)
	})
}