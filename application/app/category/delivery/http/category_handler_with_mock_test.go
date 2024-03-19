package http_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpDelivery "go-fiber-clean-architecture/application/app/category/delivery/http"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/domain/mocks"
	"go-fiber-clean-architecture/application/exception"
	"go-fiber-clean-architecture/application/helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
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

func TestGetById(t *testing.T) {
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