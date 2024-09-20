package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/product/mocks"
	"go-fiber-clean-architecture/application/app/product/usecase"
	"go-fiber-clean-architecture/application/domain"
	"testing"
)

func TestGetAll(t *testing.T) {
	// buat mock object
	mockProductRepo := new(mocks.ProductRepository)
	// buat mock data
	mockProduct := domain.Product{
		ID: 1,
		Name: "Test",
		CategoryId: 1,
	}

	// buat list mock data
	mockListProduct := make([]domain.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct)

	// testing GetAll success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetAll dari ProductRepository
		// di on method ini kita akan memanggil method GetAll dari ProductRepository, dan menambahkan parameternya context yang dibutuhkan dan return mockListProduct dan nil
		mockProductRepo.On("GetAll", mock.Anything).Return(mockListProduct, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewProductUseCase dengan parameter mockProductRepo
		productUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
		// kita panggil method GetAll dari categoryUcase
		res, err := productUcase.GetAll(context.TODO())
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockListProduct bahwa hasilnya sama
		assert.Equal(t, mockListProduct, res)
		// kita akan memastikan bahwa mockProductRepo sudah dipanggil sesuai dengan ekspektasi
		mockProductRepo.AssertExpectations(t)
	})

	// testing GetAll failed
	t.Run("error-failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetAll dari CategoryRepository
		// di on method ini kita akan memanggil method GetAll dari CategoryRepository, dan menambahkan parameternya context yang dibutuhkan dan return nil dan error
		mockProductRepo.On("GetAll", mock.Anything).Return(nil, assert.AnError).Once()
		// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
		categoryUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
		// kita panggil method GetAll dari categoryUcase
		res, err := categoryUcase.GetAll(context.TODO())
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockListCategory bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
		mockProductRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	// buat mock object
	mockProductRepo := new(mocks.ProductRepository)
	// buat mock data
	mockProduct := domain.Product{
		ID: 1,
		Name: "Test",
		CategoryId: 1,
	}

	// testing GetByID success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari ProductRepository
		// di on method ini kita akan memanggil method GetByID dari ProductRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return mockProduct dan nil
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockProduct, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewProductUseCase dengan parameter mockProductRepo
		categoryUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
		// kita panggil method GetByID dari categoryUcase
		res, err := categoryUcase.GetByID(context.TODO(), mockProduct.ID)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockProduct bahwa hasilnya sama
		assert.Equal(t, mockProduct, res)
		// kita akan memastikan bahwa mockProductRepo sudah dipanggil sesuai dengan ekspektasi
		mockProductRepo.AssertExpectations(t)
	})

	// testing GetByID failed
	t.Run("error-failed", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method GetByID dari ProductRepository
		// di on method ini kita akan memanggil method GetByID dari ProductRepository dan parameternya context(mock.anything), id(mock.AnythingOfType("int64")) yang dibutuhkan dan return nil dan error
		mockProductRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, assert.AnError).Once()
		// kita membuat object categoryUcase yang menggunakan NewProductUseCase dengan parameter mockProductRepo
		categoryUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
		// kita panggil method GetByID dari categoryUcase
		res, err := categoryUcase.GetByID(context.TODO(), mockProduct.ID)
		// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
		assert.Error(t, err)
		// kita akan membandingkan hasil res dengan mockProduct bahwa hasilnya kosong
		assert.Empty(t, res)
		// kita akan memastikan bahwa mockProductRepo sudah dipanggil sesuai dengan ekspektasi
		mockProductRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	// buat mock object
	mockProductRepo := new(mocks.ProductRepository)
	// buat mock data
	mockProduct := domain.Product{
		Name: "Test",
		CategoryId: 1,
	}

	// testing Create success
	t.Run("success", func(t *testing.T) {
		// kita akan membuat mock object untuk memanggil method Create dari ProductRepository
		// di on method ini kita akan memanggil method Create dari ProductRepository dan parameternya context(mock.anything), category(mock.AnythingOfType("domain.Product")) yang dibutuhkan dan return mockProduct dan nil
		mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Product")).Return(mockProduct, nil).Once()
		// kita membuat object categoryUcase yang menggunakan NewProductUseCase dengan parameter mockProductRepo
		productUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
		// kita panggil method Create dari productUcase
		res, err := productUcase.Create(context.TODO(), mockProduct)
		// kita akan melakukan assertion bahwa errornya nil atau bisa disebut tidak ada error
		assert.NoError(t, err)
		// kita akan membandingkan hasil res dengan mockProduct bahwa hasilnya sama
		assert.Equal(t, mockProduct, res)
		// kita akan memastikan bahwa mockProductRepo sudah dipanggil sesuai dengan ekspektasi
		mockProductRepo.AssertExpectations(t)
	})

	// testing Create failed
	//t.Run("error-failed", func(t *testing.T) {
	//	// kita akan membuat mock object untuk memanggil method Create dari CategoryRepository
	//	// di on method ini kita akan memanggil method Create dari CategoryRepository dan parameternya context(mock.anything), product(mock.AnythingOfType("domain.Category")) yang dibutuhkan dan return nil dan error
	//	mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Product")).Return(nil, assert.AnError).Once()
	//	// kita membuat object categoryUcase yang menggunakan NewCategoryUseCase dengan parameter mockCategoryRepo
	//	productUcase := usecase.NewProductUseCase(mockProductRepo, mockProductRepo)
	//	// kita panggil method Create dari productUcase
	//	res, err := productUcase.Create(context.TODO(), mockProduct)
	//	// kita akan melakukan assertion bahwa errornya tidak nil atau bisa disebut ada error
	//	assert.Error(t, err)
	//	// kita akan membandingkan hasil res dengan mockCategory bahwa hasilnya kosong
	//	assert.Empty(t, res)
	//	// kita akan memastikan bahwa mockCategoryRepo sudah dipanggil sesuai dengan ekspektasi
	//	mockProductRepo.AssertExpectations(t)
	//})
}