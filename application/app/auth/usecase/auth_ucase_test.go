package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/category/usecase"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/domain/mocks"
	"testing"
)

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
