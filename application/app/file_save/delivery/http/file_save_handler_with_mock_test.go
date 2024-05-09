package http_test

import (
	"bytes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpDelivery "go-fiber-clean-architecture/application/app/file_save/delivery/http"
	"go-fiber-clean-architecture/application/app/file_save/mocks"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"go-fiber-clean-architecture/application/helper"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllFileSaveWithMock(t *testing.T) {
	mockFileSaveUseCase := new(mocks.FileSaveUsecase)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}
	// buat list mock data
	mockListFileSave := make([]domain.FileSave, 0)
	mockListFileSave = append(mockListFileSave, mockFileSave)

	t.Run("success", func(t *testing.T) {
		mockFileSaveUseCase.On("GetAll", mock.Anything).Return(mockListFileSave, nil).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
		}


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/file-save", nil)


		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetAll},
			Method: fiber.MethodGet,
			Path: "/file-save",
		})


		// panggil handler
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})
}

func TestGetByIdFileSaveWithMock(t *testing.T) {
	mockFileSaveUseCase := new(mocks.FileSaveUsecase)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}

	t.Run("success", func(t *testing.T) {
		mockFileSaveUseCase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetByID},
			Method: fiber.MethodGet,
			Path: "/file-save/:id",
		})


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/file-save/1", nil)


		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		mockFileSaveUseCase.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockFileSaveUseCase.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, exception.ErrNotFound).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
		}
		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.GetByID},
			Method: fiber.MethodGet,
			Path: "/file-save/:id",
		})


		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/file-save/1", nil)


		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})
}

func TestCreateFileSaveWithMock(t *testing.T) {
	mockFileSaveUseCase := new(mocks.FileSaveUsecase)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}

	validate := validator.New()

	t.Run("success", func(t *testing.T) {
		mockFileSaveUseCase.On("Create", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(mockFileSave, nil).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
			Validate: validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{func(ctx *fiber.Ctx) error {
				helper.NewCustomValidation(validate, ctx)
				return ctx.Next()
			}, handler.Create},
			Method: fiber.MethodPost,
			Path: "/file-save",
		})

		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		// kita akan membuat request dengan method post dan path "/" dan bodyReq
		req := httptest.NewRequest(fiber.MethodPost, "/file-save", bodyReq)
		// kita akan set header Content-Type dengan writer.FormDataContentType()
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})

	t.Run("bad-request", func(t *testing.T) {
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
			Validate: validate,
		}

		// buat context
		app := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{func(ctx *fiber.Ctx) error {
				helper.NewCustomValidation(validate, ctx)
				return ctx.Next()
			}, handler.Create},
			Method: fiber.MethodPost,
			Path: "/file-save",
		})

		// request body
		data := strings.NewReader(`{"File": "Test"}`)
		// buat request
		req := httptest.NewRequest(fiber. MethodPost, "/file-save", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := app.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})
}

func TestUpdateFileSaveWithMock(t *testing.T) {
	mockFileSaveUseCase := new(mocks.FileSaveUsecase)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}

	validate := validator.New()

	t.Run("success", func(t *testing.T) {
		mockFileSaveUseCase.On("Update", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(mockFileSave, nil).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
			Validate: validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{func(ctx *fiber.Ctx) error {
				helper.NewCustomValidation(validate, ctx)
				return ctx.Next()
			}, handler.Update},
			Method: fiber.MethodPost,
			Path: "/file-save/:id",
		})

		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/file-save/1", bodyReq)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})

	t.Run("bad-request", func(t *testing.T) {
		mockFileSaveUseCase.On("Update", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(nil, assert.AnError).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
			Validate:        validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{func(ctx *fiber.Ctx) error {
				helper.NewCustomValidation(validate, ctx)
				return ctx.Next()
			}, handler.Update},
			Method:   fiber.MethodPost,
			Path:     "/file-save/:id",
		})

		data := strings.NewReader(`{"ID": 1, "Name": ""}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/file-save/1", data)
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

func TestDeleteFileSaveWithMock(t *testing.T) {
	mockFileSaveUseCase := new(mocks.FileSaveUsecase)

	t.Run("success", func(t *testing.T) {
		mockFileSaveUseCase.On("Delete", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("int64")).Return(nil).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Delete},
			Method: fiber.MethodDelete,
			Path: "/file-save/:id",
		})

		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/file-save/1", nil)

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockFileSaveUseCase.On("Delete", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("int64")).Return(exception.ErrNotFound).Once()
		// buat handler
		handler := httpDelivery.FileSaveHandler{
			FileSaveUseCase: mockFileSaveUseCase,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Delete},
			Method: fiber.MethodDelete,
			Path: "/file-save/:id",
		})

		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/file-save/1", nil)

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockFileSaveUseCase.AssertExpectations(t)
	})
}