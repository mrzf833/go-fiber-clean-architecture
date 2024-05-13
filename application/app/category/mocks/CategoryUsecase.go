package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
	"io"
)

type CategoryUsecase struct {
	mock.Mock
}

func (m *CategoryUsecase) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	ret := m.Called(ctx, id)
	var r0 domain.Category
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}
	return r0, ret.Error(1)
}

func (m *CategoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {
	ret := m.Called(ctx)
	var r0 []domain.Category
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.Category)
	}
	return r0, ret.Error(1)
}

func (m *CategoryUsecase) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	ret := m.Called(ctx, category)
	var r0 domain.Category
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}
	return r0, ret.Error(1)
}

func (m *CategoryUsecase) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	ret := m.Called(ctx, category)
	var r0 domain.Category
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}
	return r0, ret.Error(1)
}

func (m *CategoryUsecase) Delete(ctx context.Context, id int64) error {
	ret := m.Called(ctx, id)
	return ret.Error(0)
}

func (m *CategoryUsecase) CreateWithCsv(ctx context.Context, file io.Reader, idTrackerCategory int64) {
	m.MethodCalled("CreateWithCsv", ctx, file, idTrackerCategory)
}