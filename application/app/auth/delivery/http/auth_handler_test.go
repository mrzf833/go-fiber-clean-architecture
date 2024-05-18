package http_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-fiber-clean-architecture/application/app"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var appRun *fiber.App
var username = "john"
var password = "doe"

func TestMain(m *testing.M) {
	// run test
	fmt.Println("Before Test")
	// setup app
	appRun = app.AppInit()

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := domain.User{
		Username: username,
		Password: string(passwordHash),
	}
	// simpan data
	config.DB.Create(&user)

	tokenCh := make(chan string)
	go helper.LoginAuth(appRun, tokenCh, username, password)
	<-tokenCh

	// run test
	m.Run()

	defer config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Auth{})
	// clean data
	fmt.Println("After Test")
}

func TestAuthLogin(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat request
		reqBody := strings.NewReader(`{"Username": "` + username + `","Password": "` + password + `"}`)
		req := httptest.NewRequest(fiber.MethodPost, "/api/login", reqBody)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// bad request test
	t.Run("bad-request", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/login", nil)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func TestAuthUser(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/user", nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// Unauthorized test
	t.Run("unauthorized", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/user", nil)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}

func TestAuthLogout(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/logout", nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// Unauthorized test
	t.Run("unauthorized", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/logout", nil)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}
