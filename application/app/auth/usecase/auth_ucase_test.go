package usecase_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/auth/mocks"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/app/auth/usecase"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockAuthRepo := new(mocks.AuthRepository)
	// buat mock data
	mockAuth := domain.Auth{
		ID: 1,
		Username: "john",
		Password: "doe",
	}
	authCreateRequest := request.AuthCreateRequest{
		Username: "john",
		Password: "doe",
	}

	// testing login success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari AuthRepository
		// di on method ini kita akan memanggil method GetByUsername dari AuthRepository dan parameternya context(mock.anything), username(mock.AnythingOfType("string")) yang dibutuhkan dan return mockAuth dan nil
		mockAuthRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(mockAuth, nil).Once()
		app.Use(func(c *fiber.Ctx) error {
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo)
			// kita panggil method GetByID dari authUcase
			res, err := authUcase.Login(c, authCreateRequest)
			// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
			assert.NoError(t, err)
			// kita akan membandingkan hasil res dengan mockAuth bahwa hasilnya sama
			assert.Equal(t, mockAuth.Username, res["username"].(string))
			// kita akan memastikan bahwa mockAuthRepo sudah dipanggil sesuai dengan ekspektasi
			mockAuthRepo.AssertExpectations(t)
			return nil
		})
	})

	// testing login failed
	t.Run("failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari AuthRepository
		// di on method ini kita akan memanggil method GetByID dari AuthRepository dan parameternya context(mock.anything), username(mock.AnythingOfType("string")) yang dibutuhkan dan return nil dan error
		mockAuthRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.Auth{}, assert.AnError).Once()
		app.Use(func(c *fiber.Ctx) error {
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo)
			// kita panggil method GetByID dari authUcase
			res, err := authUcase.Login(c, authCreateRequest)
			// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
			assert.Error(t, err)
			// kita akan membandingkan hasil res dengan mockAuth bahwa hasilnya kosong
			assert.Empty(t, res)
			// kita akan memastikan bahwa mockAuthRepo sudah dipanggil sesuai dengan ekspektasi
			mockAuthRepo.AssertExpectations(t)
			return nil
		})
	})
}

func TestUser(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockAuthRepo := new(mocks.AuthRepository)
	// buat mock data
	mockAuth := domain.Auth{
		ID:       1,
		Username: "johns",
		Password: "doe",
	}

	var res jwt.MapClaims
	// testing login success
	t.Run("success", func(t *testing.T) {
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("user", jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": mockAuth.Username,
			}))
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo)
			// kita panggil method GetByID dari authUcase
			res = authUcase.User(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		assert.Equal(t, mockAuth.Username, res["username"].(string))
		mockAuthRepo.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockAuthRepo := new(mocks.AuthRepository)
	// buat mock data
	mockAuth := domain.Auth{
		ID:       1,
		Username: "johns",
		Password: "doe",
	}
	helper.ConnectRedis()
	// testing login success
	t.Run("success", func(t *testing.T) {
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("user", jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": mockAuth.Username,
			}))
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo)
			// kita panggil method GetByID dari authUcase
			authUcase.Logout(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		mockAuthRepo.AssertExpectations(t)
	})
}


