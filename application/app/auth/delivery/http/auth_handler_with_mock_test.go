package http_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpDelivery "go-fiber-clean-architecture/application/app/auth/delivery/http"
	"go-fiber-clean-architecture/application/app/auth/usecase"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/domain/mocks"
	"go-fiber-clean-architecture/application/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAuthLoginWithMock(t *testing.T) {
	returnAuthUseCase := map[string]interface{}{
		"token": "token",
		"expire": time.Now().Add(config.ExpireToken).Unix(),
		"username": "test",
	}

	validate := validator.New()
	t.Run("success", func(t *testing.T) {
		mockAuthUseCase := new(mocks.AuthUsecase)
		mockAuthUseCase.On("Login", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("request.AuthCreateRequest")).Return(returnAuthUseCase, nil).Once()


		// buat handler
		handler := httpDelivery.AuthHandler{
			Validate: validate,
			Ucase: mockAuthUseCase,
		}
		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Login},
			Method: fiber.MethodPost,
			Path: "/login",
		})

		data := strings.NewReader(`{"Username": "test","Password": "test"}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/login", data)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler untuk di testing
		res, err := c.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		//bytes, err := io.ReadAll(res.Body)
		//println(string(bytes))
		mockAuthUseCase.AssertExpectations(t)
	})

	t.Run("bad-request", func(t *testing.T) {
		mockAuthRepository := new(mocks.AuthRepository)
		mockAuthRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.Auth{}, assert.AnError).Once()
		mockAuthUseCase := usecase.NewAuthUseCase(mockAuthRepository)
		//mockAuthUseCase.On("Login", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("request.AuthCreateRequest")).Return(map[string]interface{}{}, assert.AnError).Once()
		// buat handler
		handler := httpDelivery.AuthHandler{
			Ucase: mockAuthUseCase,
			Validate:        validate,
		}

		// buat context
		c := helper.TestApp(helper.HelperRouter{
			Handlers: []fiber.Handler{handler.Login},
			Method:   fiber.MethodPost,
			Path:     "/login",
		})

		data := strings.NewReader(`{"Username": "test","Password": "test"}`)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/login", data)
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

func TestAuthUserWithMock(t *testing.T) {
	validate := validator.New()
	mockAuthUseCase := new(mocks.AuthUsecase)
	mockAuthUseCase.On("User", mock.AnythingOfType("*fiber.Ctx")).Return(jwt.MapClaims{"username": "test"}, nil).Once()
	// buat handler
	handler := httpDelivery.AuthHandler{
		Validate: validate,
		Ucase: mockAuthUseCase,
	}
	// buat context
	c := helper.TestApp(helper.HelperRouter{
		Handlers: []fiber.Handler{handler.User},
		Method: fiber.MethodGet,
		Path: "/user",
	})

	// buat request
	req := httptest.NewRequest(fiber.MethodGet, "/user", nil)
	req.Header.Set("Content-Type", "application/json")

	// panggil handler untuk di testing
	res, err := c.Test(req)
	// cek error
	assert.NoError(t, err)
	// cek response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	println(string(bytes))
	mockAuthUseCase.AssertExpectations(t)
}