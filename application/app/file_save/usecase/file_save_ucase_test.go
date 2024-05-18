package usecase_test

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/file_save/mocks"
	"go-fiber-clean-architecture/application/app/file_save/usecase"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAll(t *testing.T) {
	// buat mock object
	mockFileSaveRepo := new(mocks.FileSaveRepository)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}

	// buat list mock data
	mockListFileSave := make([]domain.FileSave, 0)
	mockListFileSave = append(mockListFileSave, mockFileSave)

	// testing GetAll success
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method GetAll dari FileSaveRepository
		// di on method ini kita akan memanggil method GetAll dari FileSaveRepository, dan menambahkan parameternya context yang dibutuhkan dan return mockListFileSave dan nil
		mockFileSaveRepo.On("GetAll", mock.AnythingOfType("*fiber.Ctx")).Return(mockListFileSave, nil).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)
		// kita panggil method GetAll dari fileSaveUcase

		var err error
		var res []domain.FileSave
		app.Use(func(c *fiber.Ctx) error {
			res, err = fileSaveUcase.GetAll(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockListFileSave bahwa hasilnya sama
		assert.Equal(t, mockListFileSave, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})

	// testing GetAll failed
	t.Run("error-failed", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method GetAll dari FileSaveRepository
		// di on method ini kita akan memanggil method GetAll dari FileSaveRepository, dan menambahkan parameternya context yang dibutuhkan dan return nil dan error
		mockFileSaveRepo.On("GetAll", mock.AnythingOfType("*fiber.Ctx")).Return(nil, assert.AnError).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var err error
		var res []domain.FileSave
		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method GetAll dari fileSaveUcase
			res, err = fileSaveUcase.GetAll(c)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockListFileSave bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	// buat mock object
	mockFileSaveRepo := new(mocks.FileSaveRepository)
	// buat mock data
	mockFileSave := domain.FileSave{
		ID: 1,
		Name: "Test",
	}

	// testing GetByID success
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("GetByID", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var err error
		var res domain.FileSave
		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method GetByID dari fileSaveUcase
			res, err = fileSaveUcase.GetByID(c, mockFileSave.ID)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya sama
		assert.Equal(t, mockFileSave, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})

	// testing GetByID failed
	t.Run("error-failed", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return nil dan error
		mockFileSaveRepo.On("GetByID", mock.AnythingOfType("*fiber.Ctx"), mock.AnythingOfType("int64")).Return(nil, assert.AnError).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var err error
		var res domain.FileSave
		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method GetByID dari fileSaveUcase
			res, err = fileSaveUcase.GetByID(c, mockFileSave.ID)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodGet, "/", nil)
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	// buat mock object
	mockFileSaveRepo := new(mocks.FileSaveRepository)
	// buat mock data
	mockFileSave := domain.FileSave{
		Name: "Test",
	}

	// testing Create success
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method Create dari FileSaveRepository
		// di on method ini kita akan memanggil method Create dari FileSaveRepository dan parameternya context(mock.anything), fileSave(mock.AnythingOfType("domain.FileSave")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(mockFileSave, nil).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var err error
		var res domain.FileSave

		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Create dari fileSaveUcase
			res, err = fileSaveUcase.Create(c, mockFileSave)
			return nil
		})
		bodyReq := new(bytes.Buffer)
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()
		req := httptest.NewRequest(fiber.MethodPost, "/", bodyReq)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya sama

		applicationPath := utils.GetApplicationPath()
		err = os.Remove(applicationPath + "/storage/public/" + res.Name)
		assert.NoError(t, err)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})

	// testing Create failed
	t.Run("error-failed", func(t *testing.T) {
		app := fiber.New()
		// kita akan membuat mock object untuk memanggil method Create dari FileSaveRepository
		// di on method ini kita akan memanggil method Create dari FileSaveRepository dan parameternya context(mock.anything), fileSave(mock.AnythingOfType("domain.FileSave")) yang dibutuhkan dan return nil dan error
		mockFileSaveRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(nil, assert.AnError).Once()
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var err error
		var res domain.FileSave

		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Create dari fileSaveUcase
			res, err = fileSaveUcase.Create(c, mockFileSave)
			return nil
		})
		bodyReq := new(bytes.Buffer)
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		req := httptest.NewRequest(fiber.MethodPost, "/", bodyReq)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	// buat mock object
	mockFileSaveRepo := new(mocks.FileSaveRepository)

	//testing Update success
	t.Run("success", func(t *testing.T) {
		app := fiber.New()

		// buat mock data
		mockFileSave := domain.FileSave{
			ID: 1,
			Name: "",
		}

		storagePublicPath := utils.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		mockFileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		db, mockDb := utils.NewMockDB()
		// ----------------------------- kenapa ini ada 3 on method? -----------------------------
		// karena di update method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID dan GetDb
		// baru kemudian kita melakukan update data menggunakan method Update

		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()

		// kita akan membuat mock object untuk memanggil method GetDb dari FileSaveRepository
		// karena ini memakai transaction, maka kita perlu memanggil exceptionBegin dan exceptionCommit
		mockDb.ExpectBegin()
		mockFileSaveRepo.On("GetDb").Return(db).Once()
		mockDb.ExpectCommit()


		// kita akan membuat mock object untuk memanggil method Update dari FileSaveRepository
		// di on method ini kita akan memanggil method Update dari FileSaveRepository dan parameternya context(mock.anything), fileSave(mock.AnythingOfType("domain.FileSave")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(mockFileSave, nil).Once()
		// ------------------------------------------------
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var res domain.FileSave
		// kita akan membuat mock object untuk memanggil method Update dari FileSaveRepository
		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Update dari fileSaveUcase
			res, err = fileSaveUcase.Update(c, mockFileSave)
			return nil
		})

		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		// kita akan membuat request dengan method post dan path "/" dan bodyReq
		req := httptest.NewRequest(fiber.MethodPost, "/", bodyReq)
		// kita akan set header Content-Type dengan writer.FormDataContentType()
		req.Header.Set("Content-Type", writer.FormDataContentType())
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)

		defer func() {
			err = os.Remove(storagePublicPath + res.Name)
			assert.NoError(t, err)
		}()
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya sama
		//assert.Equal(t, mockFileSave, res)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})

	// testing Update failed
	t.Run("error-failed", func(t *testing.T) {
		app := fiber.New()

		// buat mock data
		mockFileSave := domain.FileSave{
			ID: 1,
			Name: "",
		}

		storagePublicPath := utils.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		mockFileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		db, mockDb := utils.NewMockDB()
		// ----------------------------- kenapa ini ada 3 on method? -----------------------------
		// karena di update method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID dan GetDb
		// baru kemudian kita melakukan update data menggunakan method Update


		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()

		// kita akan membuat mock object untuk memanggil method GetDb dari FileSaveRepository
		// karena ini memakai transaction, maka kita perlu memanggil exceptionBegin dan exceptionCommit
		mockDb.ExpectBegin()
		mockFileSaveRepo.On("GetDb").Return(db).Once()
		mockDb.ExpectCommit()


		// kita akan membuat mock object untuk memanggil method Update dari FileSaveRepository
		// di on method ini kita akan memanggil method Update dari FileSaveRepository dan parameternya context(mock.anything), fileSave(mock.AnythingOfType("domain.FileSave")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.FileSave")).Return(nil, assert.AnError).Once()
		// ------------------------------------------------
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		var res domain.FileSave
		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Update dari fileSaveUcase
			res, err = fileSaveUcase.Update(c, mockFileSave)
			return nil
		})

		// kita akan membuat file baru
		bodyReq := new(bytes.Buffer)
		// kita akan membuat writer untuk membuat form file
		writer := multipart.NewWriter(bodyReq)
		part, _ := writer.CreateFormFile("file", "test.txt")
		part.Write([]byte("test"))
		writer.Close()

		// kita akan membuat request dengan method post dan path "/" dan bodyReq
		req := httptest.NewRequest(fiber.MethodPost, "/", bodyReq)
		// kita akan set header Content-Type dengan writer.FormDataContentType()
		req.Header.Set("Content-Type", writer.FormDataContentType())
		app.Test(req)

		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockFileSave bahwa hasilnya kosong
		assert.Empty(t, res)

		defer func() {
			err = os.Remove(storagePublicPath + mockFileSave.Name)
			assert.NoError(t, err)
		}()
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	// buat mock object
	mockFileSaveRepo := new(mocks.FileSaveRepository)
	db, mockDb := utils.NewMockDB()
	// buat mock data

	// testing Delete success
	t.Run("success", func(t *testing.T) {
		app := fiber.New()

		// buat mock data
		mockFileSave := domain.FileSave{
			ID: 1,
			Name: "",
		}

		storagePublicPath := utils.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		mockFileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		// ----------------------------- kenapa ini ada 3 on method? -----------------------------
		// karena di delete method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID and GetDb
		// baru kemudian kita melakukan delete data menggunakan method Delete

		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()

		// kita akan membuat mock object untuk memanggil method GetDb dari FileSaveRepository
		// karena ini memakai transaction, maka kita perlu memanggil exceptionBegin dan exceptionCommit
		mockDb.ExpectBegin()
		mockFileSaveRepo.On("GetDb").Return(db).Once()
		mockDb.ExpectCommit()

		// kita akan membuat mock object untuk memanggil method Delete dari FileSaveRepository
		// di on method ini kita akan memanggil method Delete dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return nil
		mockFileSaveRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		// ------------------------------------------------
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Delete dari fileSaveUcase
			err = fileSaveUcase.Delete(c, mockFileSave.ID)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodDelete, "/", nil)
		app.Test(req)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})

	// testing Delete failed
	t.Run("error-failed", func(t *testing.T) {
		app := fiber.New()

		// buat mock data
		mockFileSave := domain.FileSave{
			ID: 1,
			Name: "",
		}

		storagePublicPath := utils.GetStoragePublicPath()
		temp, err := os.CreateTemp(storagePublicPath+"/upload_file", "")
		assert.NoError(t, err)

		mockFileSave.Name = "/upload_file/" +  filepath.Base(temp.Name())
		// ----------------------------- kenapa ini ada 3 on method? -----------------------------
		// karena di delete method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID
		// baru kemudian kita melakukan delete data menggunakan method Delete

		// kita akan membuat mock object untuk memanggil method GetByID dari FileSaveRepository
		// di on method ini kita akan memanggil method GetByID dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockFileSave dan nil
		mockFileSaveRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockFileSave, nil).Once()

		// kita akan membuat mock object untuk memanggil method GetDb dari FileSaveRepository
		// karena ini memakai transaction, maka kita perlu memanggil exceptionBegin dan exceptionCommit
		mockDb.ExpectBegin()
		mockFileSaveRepo.On("GetDb").Return(db).Once()
		mockDb.ExpectCommit()


		// kita akan membuat mock object untuk memanggil method Delete dari FileSaveRepository
		// di on method ini kita akan memanggil method Delete dari FileSaveRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return assert.AnError
		mockFileSaveRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(assert.AnError).Once()
		// ------------------------------------------------
		// kita membuat object fileSaveUcase yang menggunakan NewFileSaveUseCase dengan parameter mockFileSaveRepo
		fileSaveUcase := usecase.NewFileSaveUseCase(mockFileSaveRepo)

		app.Use(func(c *fiber.Ctx) error {
			// kita panggil method Delete dari fileSaveUcase
			err = fileSaveUcase.Delete(c, mockFileSave.ID)
			return nil
		})
		req := httptest.NewRequest(fiber.MethodDelete, "/", nil)
		app.Test(req)

		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)

		defer func() {
			err = os.Remove(storagePublicPath + mockFileSave.Name)
			assert.NoError(t, err)
		}()
		// kita akan memastikan bahwa mockFileSaveRepo sudah dipanggil sesuai dengan ekspektasi
		mockFileSaveRepo.AssertExpectations(t)
	})
}