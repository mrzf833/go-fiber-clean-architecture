package http_test

import (
	"bytes"
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
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
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
	user := domain.User{
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

func TestGetAllFileSave(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat data terlebih dahulu
		fileSave := domain.FileSave{
			Name: "file 1",
		}
		// simpan data
		config.DB.Create(&fileSave)

		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/file-save", nil)
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
		config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.FileSave{})
	})
}

func TestGetByIdFileSave(t *testing.T) {
	// buat data terlebih dahulu
	fileSave := domain.FileSave{
		Name: "FileSave 1",
	}
	// simpan data
	config.DB.Create(&fileSave)

	// success test
	t.Run("success", func(t *testing.T) {
		// get id
		id := strconv.FormatInt(fileSave.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/file-save/" + id, nil)
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
		id := strconv.FormatInt(fileSave.ID + 1, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodGet, "/api/file-save/" + id, nil)
		req.Header.Set("Authorization", "Bearer " + helper.Token)

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	// hapus data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.FileSave{})
}

func TestCreateFileSave(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPost, "/api/file-save", bodyReq)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		defer func() {

			// get data response body bytes
			bytesBody, _ := io.ReadAll(res.Body)
			var resBody map[string]interface{}
			// unmarshal bytes body to map
			json.Unmarshal(bytesBody, &resBody)
			// get data response body data
			resData := resBody["data"].(map[string]interface{})
			storagePublicPath := helper.GetStoragePublicPath()
			err = os.Remove(storagePublicPath + resData["name"].(string))
			assert.NoError(t, err)
		}()
	})

	// bad request test
	t.Run("bad-request", func(t *testing.T) {
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/file-save", nil)
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
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.FileSave{})
}

func TestUpdateFileSave(t *testing.T) {
	// success test
	t.Run("success", func(t *testing.T) {
		// buat data terlebih dahulu
		fileSave := domain.FileSave{
			Name: "",
		}

		storagePublicPath := helper.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		fileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		// simpan data
		config.DB.Create(&fileSave)
		// get id
		id := strconv.FormatInt(fileSave.ID, 10)
		// buat request
		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPost, "/api/file-save/" + id, bodyReq)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", writer.FormDataContentType())

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
		assert.NotEqual(t, fileSave.Name, dataJson["data"].(map[string]interface{})["name"])

		defer func() {
			err = os.Remove(storagePublicPath + dataJson["data"].(map[string]interface{})["name"].(string))
			assert.NoError(t, err)
		}()
	})

	// bad request test
	t.Run("bad-request", func(t *testing.T) {
		fileSave := domain.FileSave{
			Name: "",
		}
		// simpan data
		config.DB.Create(&fileSave)
		// get id
		id := strconv.FormatInt(fileSave.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodPost, "/api/file-save/" + id, nil)
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
		fileSave := domain.FileSave{
			Name: "",
		}
		// simpan data
		config.DB.Create(&fileSave)
		// get id
		id := strconv.FormatInt(fileSave.ID + 1, 10)
		// buat request
		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPost, "/api/file-save/" + id, bodyReq)
		req.Header.Set("Authorization", "Bearer " + helper.Token)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// panggil handler
		res, err := appRun.Test(req)
		// cek error
		assert.NoError(t, err)
		// cek response
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	// delete data
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.FileSave{})
}

func TestDeleteFileSave(t *testing.T) {
	// buat data terlebih dahulu

	// success test
	t.Run("success", func(t *testing.T) {
		fileSave := domain.FileSave{
			Name: "",
		}
		storagePublicPath := helper.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		fileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		// simpan data
		config.DB.Create(&fileSave)
		// get id
		id := strconv.FormatInt(fileSave.ID, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/api/file-save/" + id, nil)
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
		fileSave := domain.FileSave{
			Name: "",
		}
		// simpan data
		config.DB.Create(&fileSave)
		// get id
		id := strconv.FormatInt(fileSave.ID + 1, 10)
		// buat request
		req := httptest.NewRequest(fiber.MethodDelete, "/api/file-save/" + id, nil)
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
	config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.FileSave{})
}