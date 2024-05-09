package mocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
)

type FileSaveUsecase struct {
	mock.Mock
}

func (m *FileSaveUsecase) GetByID(ctx *fiber.Ctx, id int64) (domain.FileSave, error) {
	ret := m.Called(ctx, id)
	var r0 domain.FileSave
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
	}
	return r0, ret.Error(1)
}

func (m *FileSaveUsecase) GetAll(c *fiber.Ctx) ([]domain.FileSave, error) {
	ret := m.Called(c)
	var r0 []domain.FileSave
	if ret.Get(1) == nil {
		r0 = ret.Get(0).([]domain.FileSave)
	}
	return r0, ret.Error(1)
}

func (m *FileSaveUsecase) Create(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	ret := m.Called(c, fileSave)
	var r0 domain.FileSave
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
	}
	return r0, ret.Error(1)
}

func (m *FileSaveUsecase) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	ret := m.Called(c, fileSave)
	var r0 domain.FileSave
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.FileSave)
	}
	return r0, ret.Error(1)
}

func (m *FileSaveUsecase) Delete(c *fiber.Ctx, id int64) error {
	ret := m.Called(c, id)
	return ret.Error(0)
}