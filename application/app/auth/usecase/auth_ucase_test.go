package usecase_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocksAuth "go-fiber-clean-architecture/application/app/auth/mocks"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/app/auth/usecase"
	mocksUser "go-fiber-clean-architecture/application/app/user/mocks"
	"go-fiber-clean-architecture/application/domain"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockUserRepo := new(mocksUser.UserRepository)
	mockAuthRepo := new(mocksAuth.AuthRepository)
	// buat mock data
	mockUser := domain.User{
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
		// kita akan membuat mock object untuk memanggil method GetByID dari UserRepository
		// di on method ini kita akan memanggil method GetByUsername dari UserRepository dan parameternya context(mock.anything), username(mock.AnythingOfType("string")) yang dibutuhkan dan return mockAuth dan nil
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		mockAuthRepo.On("CreateToken", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		app.Use(func(c *fiber.Ctx) error {
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo, mockUserRepo)
			// kita panggil method GetByID dari authUcase
			res, err := authUcase.Login(c, authCreateRequest)
			// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
			assert.NoError(t, err)
			// kita akan membandingkan hasil res dengan mockAuth bahwa hasilnya sama
			assert.Equal(t, mockUser.Username, res["username"].(string))
			// kita akan memastikan bahwa mockAuthRepo sudah dipanggil sesuai dengan ekspektasi
			mockUserRepo.AssertExpectations(t)
			mockAuthRepo.AssertExpectations(t)
			return nil
		})
	})

	// testing login failed
	t.Run("failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari AuthRepository
		// di on method ini kita akan memanggil method GetByID dari UserRepository dan parameternya context(mock.anything), username(mock.AnythingOfType("string")) yang dibutuhkan dan return nil dan error
		mockAuthRepo.On("CreateToken", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
		app.Use(func(c *fiber.Ctx) error {
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo dan mockUserRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo, mockUserRepo)
			// kita panggil method GetByID dari authUcase
			res, err := authUcase.Login(c, authCreateRequest)
			// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
			assert.Error(t, err)
			// kita akan membandingkan hasil res dengan mockAuth bahwa hasilnya kosong
			assert.Empty(t, res)
			// kita akan memastikan bahwa mockAuthRepo sudah dipanggil sesuai dengan ekspektasi
			mockUserRepo.AssertExpectations(t)
			mockAuthRepo.AssertExpectations(t)
			return nil
		})
	})
}

func TestUser(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockUserRepo := new(mocksUser.UserRepository)
	mockAuthRepo := new(mocksAuth.AuthRepository)
	// buat mock data
	mockUser := domain.User{
		ID:       1,
		Username: "johns",
		Password: "doe",
	}

	var res jwt.MapClaims
	// testing login success
	t.Run("success", func(t *testing.T) {
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("user", jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": mockUser.Username,
			}))
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo dan mockUserRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo, mockUserRepo)
			// kita panggil method GetByID dari authUcase
			res = authUcase.User(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		assert.Equal(t, mockUser.Username, res["username"].(string))
		mockUserRepo.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	app := fiber.New()
	// buat mock object
	mockUserRepo := new(mocksUser.UserRepository)
	mockAuthRepo := new(mocksAuth.AuthRepository)
	// buat mock data
	mockUser := domain.User{
		ID:       1,
		Username: "johns",
		Password: "doe",
	}
	// testing login success
	t.Run("success", func(t *testing.T) {
		mockAuthRepo.On("DeleteToken", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("user", jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": mockUser.Username,
			}))
			// kita membuat object authUcase yang menggunakan NewAuthUseCase dengan parameter mockAuthRepo dan mockUserRepo
			authUcase := usecase.NewAuthUseCase(mockAuthRepo, mockUserRepo)
			// kita panggil method GetByID dari authUcase
			authUcase.Logout(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		mockAuthRepo.AssertExpectations(t)
	})
}


