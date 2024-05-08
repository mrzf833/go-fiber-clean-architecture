package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/category/mocks"
	"go-fiber-clean-architecture/application/app/category/usecase"
	"go-fiber-clean-architecture/application/domain"
	"testing"
)

func TestGetAll(t *testing.T) {
	// buat mock object
	mockCategoryRepo := new(mocks.CategoryRepository)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	// buat list mock data
	mockListCategory := make([]domain.Category, 0)
	mockListCategory = append(mockListCategory, mockCategory)

	// testing GetAll success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetAll dari CategoryRepository
		// di on method ini kita akan memanggil method GetAll dari CategoryRepository, dan menambahkan parameternya context yang dibutuhkan dan return mockListCategory dan nil
		mockCategoryRepo.On("GetAll", mock.Anything).Return(mockListCategory, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method GetAll dari categoryUcase
		res, err := categoryUcase.GetAll(context.TODO())
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockListCategory bahwa hasilnya sama
		assert.Equal(t, mockListCategory, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})

	// testing GetAll failed
	t.Run("error-failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetAll dari CategoryRepository
		// di on method ini kita akan memanggil method GetAll dari CategoryRepository, dan menambahkan parameternya context yang dibutuhkan dan return nil dan error
		mockCategoryRepo.On("GetAll", mock.Anything).Return(nil, assert.AnError).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method GetAll dari categoryUcase
		res, err := categoryUcase.GetAll(context.TODO())
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockListCategory bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	// buat mock object
	mockCategoryRepo := new(mocks.CategoryRepository)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	// testing GetByID success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method GetByID dari categoryUcase
		res, err := categoryUcase.GetByID(context.TODO(), mockCategory.ID)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya sama
		assert.Equal(t, mockCategory, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})

	// testing GetByID failed
	t.Run("error-failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return nil dan error
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Category{}, assert.AnError).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method GetByID dari categoryUcase
		res, err := categoryUcase.GetByID(context.TODO(), mockCategory.ID)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	// buat mock object
	mockCategoryRepo := new(mocks.CategoryRepository)
	// buat mock data
	mockCategory := domain.Category{
		Name: "Test",
	}

	// testing Create success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method Create dari CategoryRepository
		// di on method ini kita akan memanggil method Create dari CategoryRepository dan parameternya context(mock.anything), category(mock.AnythingOfType("domain.Category")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Category")).Return(mockCategory, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Create dari categoryUcase
		res, err := categoryUcase.Create(context.TODO(), mockCategory)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya sama
		assert.Equal(t, mockCategory, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})

	// testing Create failed
	t.Run("error-failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method Create dari CategoryRepository
		// di on method ini kita akan memanggil method Create dari CategoryRepository dan parameternya context(mock.anything), category(mock.AnythingOfType("domain.Category")) yang dibutuhkan dan return nil dan error
		mockCategoryRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Category")).Return(domain.Category{}, assert.AnError).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Create dari categoryUcase
		res, err := categoryUcase.Create(context.TODO(), mockCategory)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	// buat mock object
	mockCategoryRepo := new(mocks.CategoryRepository)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	// testing Update success
	t.Run("success", func(t *testing.T) {
		// ----------------------------- kenapa ini ada 2 on method? -----------------------------
		// karena di update method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID
		// baru kemudian kita melakukan update data menggunakan method Update


		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// kita akan membuat mock object untuk memanggil method Update dari CategoryRepository
		// di on method ini kita akan memanggil method Update dari CategoryRepository dan parameternya context(mock.anything), category(mock.AnythingOfType("domain.Category")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.Category")).Return(mockCategory, nil).Once()
		// ------------------------------------------------
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Update dari categoryUcase
		res, err := categoryUcase.Update(context.TODO(), mockCategory)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya sama
		assert.Equal(t, mockCategory, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})

	// testing Update failed
	t.Run("error-failed", func(t *testing.T) {
		// ----------------------------- kenapa ini ada 2 on method? -----------------------------
		// karena di update method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID
		// baru kemudian kita melakukan update data menggunakan method Update


		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// kita akan membuat mock object untuk memanggil method Update dari CategoryRepository
		// di on method ini kita akan memanggil method Update dari CategoryRepository dan parameternya context(mock.anything), category(mock.AnythingOfType("domain.Category")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.Category")).Return(domain.Category{}, assert.AnError).Once()
		// ------------------------------------------------
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Update dari categoryUcase
		res, err := categoryUcase.Update(context.TODO(), mockCategory)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	// buat mock object
	mockCategoryRepo := new(mocks.CategoryRepository)
	// buat mock data
	mockCategory := domain.Category{
		ID: 1,
		Name: "Test",
	}

	// testing Delete success
	t.Run("success", func(t *testing.T) {
		// ----------------------------- kenapa ini ada 2 on method? -----------------------------
		// karena di delete method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID
		// baru kemudian kita melakukan delete data menggunakan method Delete

		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// kita akan membuat mock object untuk memanggil method Delete dari CategoryRepository
		// di on method ini kita akan memanggil method Delete dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return nil
		mockCategoryRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		// ------------------------------------------------
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Delete dari categoryUcase
		err := categoryUcase.Delete(context.TODO(), mockCategory.ID)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})

	// testing Delete failed
	t.Run("error-failed", func(t *testing.T) {
		// ----------------------------- kenapa ini ada 2 on method? -----------------------------
		// karena di delete method kita melakukan pengecekan apakah data ada atau tidak menggunakan method GetByID
		// baru kemudian kita melakukan delete data menggunakan method Delete

		// kita akan membuat mock object untuk memanggil method GetByID dari CategoryRepository
		// di on method ini kita akan memanggil method GetByID dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockCategory dan nil
		mockCategoryRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockCategory, nil).Once()
		// kita akan membuat mock object untuk memanggil method Delete dari CategoryRepository
		// di on method ini kita akan memanggil method Delete dari CategoryRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return assert.AnError
		mockCategoryRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(assert.AnError).Once()
		// ------------------------------------------------
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewCategoryUseCase(mockCategoryRepo)
		// kita panggil method Delete dari categoryUcase
		err := categoryUcase.Delete(context.TODO(), mockCategory.ID)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockCategoryRepo.AssertExpectations(t)
	})
}