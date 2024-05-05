package http_test

import (
	"encoding/json"
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
	"strconv"
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

	// create user
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := domain.Auth{
		Username: username,
		Password: string(passwordHash),
	}
	// simpan data
	config.DB.Create(&user)

	// setup token
	tokenCh := make(chan string)
	go helper.LoginAuth(appRun, tokenCh, "john", "doe")
	<-tokenCh

	// run test
	m.Run()

	defer config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Auth{})

	// clean data
	fmt.Println("After Test")
}

func TestGetAllCategory(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat data terlebih dahulu
		category := domain.Category{
			Name: "Category 1",
		}
		// simpan data
		config.DB.Create(&category)

		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/category", nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
		// cek data berapa banyak
		dataJson := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&dataJson)
		if err != nil {
			fmt.Println(err)
		}
		lenghtData := len(dataJson["data"].([]interface{}))
		assert.Greater(t, lenghtData, 0)

		// hapus data
		config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Category{})
	})
}

func TestGetByIdCategory(t *testing.T) {
	// buat data terlebih dahulu
	category := domain.Category{
		Name: "Category 1",
	}
	// simpan data
	config.DB.Create(&category)

	// success test
	t.Run("success", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/category/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// not found test
	t.Run("not-found", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID + 1, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/category/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	// hapus data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Category{})
}

func TestCreateCategory(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat request
		reqBody := strings.NewReader(`{"Name": "Test"}`)
		req := httptest.NewRequest(fiber.MethodPost, "/api/category", reqBody)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	// bad request test
	t.Run("bad-request", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/category", nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	// delete data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Category{})
}

func TestUpdateCategory(t *testing.T) {
	// buat data terlebih dahulu
	category := domain.Category{
		Name: "Category 1",
	}
	// simpan data
	config.DB.Create(&category)

	// success test
	t.Run("success", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID, 10)
		// buat request
		changeName := "Test Berubah"
		reqBody := strings.NewReader(`{"Name": "`+changeName+`"}`)
		req := httptest.NewRequest(fiber.MethodPut, "/api/category/" + id, reqBody)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)

		// cek data berapa banyak
		dataJson := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&dataJson)
		assert.NoError(t, err)
		assert.Equal(t, changeName, dataJson["data"].(map[string]interface{})["name"])
	})

	// bad request test
	t.Run("bad-request", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodPut, "/api/category/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	// not found test
	t.Run("not-found", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID + 1, 10)
		// buat request
		requestBody := strings.NewReader(`{"Name": "Test Berubah"}`)
		req := httptest.NewRequest(fiber.MethodPut, "/api/category/" + id, requestBody)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	// delete data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Category{})
}

func TestDeleteCategory(t *testing.T) {
	// buat data terlebih dahulu
	category := domain.Category{
		Name: "Category 1",
	}
	// simpan data
	config.DB.Create(&category)

	// success test
	t.Run("success", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/api/category/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// not found test
	t.Run("not-found", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(category.ID + 1, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/api/category/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", "application/json")

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	// delete data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Category{})
}