package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
)

type CategoryRepository struct {
	mock.Mock
}

func (m *CategoryRepository) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	// pemanggilan fungsi GetByID dengan parameter ctx dan id menggunakan mock
	ret := m.Called(ctx, id)

	// pengembalian value dari fungsi GetByID
	var r0 domain.Category

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}
	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	// pemanggilan fungsi GetAll dengan parameter ctx menggunakan mock
	ret := m.Called(ctx)

	// pengembalian value dari fungsi GetAll
	var r0 []domain.Category

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.Category)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *CategoryRepository) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	// pemanggilan fungsi Create dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, category)

	// pengembalian value dari fungsi Create
	var r0 domain.Category

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *CategoryRepository) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	// pemanggilan fungsi Update dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, category)

	// pengembalian value dari fungsi Update
	var r0 domain.Category

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Category)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *CategoryRepository) Delete(ctx context.Context, id int64) error {
	// pemanggilan fungsi Delete dengan parameter ctx dan id menggunakan mock
	ret := m.Called(ctx, id)

	// pengembalian value dari fungsi Delete
	return ret.Error(0)
}

func (m *CategoryRepository) CreateAll(ctx context.Context, category []domain.Category) ([]domain.Category, error) {
	// pemanggilan fungsi CreateAll dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, category)

	// pengembalian value dari fungsi CreateAll
	var r0 []domain.Category

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).([]domain.Category)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *CategoryRepository) CreateInBatches(ctx context.Context, category []domain.Category, size int) error {
	// pemanggilan fungsi CreateInBatches dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, category, size)

	// pengembalian value dari fungsi CreateInBatches
	return ret.Error(0)
}