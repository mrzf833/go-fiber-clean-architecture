package mocks

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
	"gorm.io/gorm"
)

type FileSaveRepository struct {
	mock.Mock
}

func (m *FileSaveRepository) GetDb() *gorm.DB {
	ret := m.Called()
	return ret.Get(0).(*gorm.DB)
}

func (m *FileSaveRepository) GetByID(c *fiber.Ctx, id int64) (domain.FileSave, error) {
	// pemanggilan fungsi GetByID dengan parameter ctx dan id menggunakan mock
	ret := m.Called(c, id)

	// pengembalian value dari fungsi GetByID
	var r0 domain.FileSave

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
	}
	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *FileSaveRepository) GetAll(c *fiber.Ctx) ([]domain.FileSave, error) {
	// pemanggilan fungsi GetAll dengan parameter ctx menggunakan mock
	ret := m.Called(c)

	// pengembalian value dari fungsi GetAll
	var r0 []domain.FileSave

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).([]domain.FileSave)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *FileSaveRepository) Create(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// pemanggilan fungsi Create dengan parameter ctx dan category menggunakan mock
	ret := m.Called(c, fileSave)

	// pengembalian value dari fungsi Create
	var r0 domain.FileSave

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
		r0.Name = fileSave.Name
	}
	fmt.Println(ret)
	fmt.Println(fileSave)
	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *FileSaveRepository) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// pemanggilan fungsi Update dengan parameter ctx dan category menggunakan mock
	ret := m.Called(c, fileSave)

	// pengembalian value dari fungsi Update
	var r0 domain.FileSave

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
		r0.Name = fileSave.Name
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *FileSaveRepository) Delete(c *fiber.Ctx, id int64) error {
	// pemanggilan fungsi Delete dengan parameter ctx dan id menggunakan mock
	ret := m.Called(c, id)

	// pengembalian value dari fungsi Delete
	return ret.Error(0)
}