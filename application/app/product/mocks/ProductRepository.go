package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
)

type ProductRepository struct {
	mock.Mock
}

func (m *ProductRepository) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	// pemanggilan fungsi GetByID dengan parameter ctx dan id menggunakan mock
	ret := m.Called(ctx, id)

	// pengembalian value dari fungsi GetByID
	var r0 domain.Product

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Product)
	}
	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	// pemanggilan fungsi GetAll dengan parameter ctx menggunakan mock
	ret := m.Called(ctx)

	// pengembalian value dari fungsi GetAll
	var r0 []domain.Product

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.Product)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *ProductRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	// pemanggilan fungsi Create dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, product)

	// pengembalian value dari fungsi Create
	var r0 domain.Product

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Product)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *ProductRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	// pemanggilan fungsi Update dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, product)

	// pengembalian value dari fungsi Update
	var r0 domain.Product

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Product)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *ProductRepository) Delete(ctx context.Context, id int64) error {
	// pemanggilan fungsi Delete dengan parameter ctx dan id menggunakan mock
	ret := m.Called(ctx, id)

	// pengembalian value dari fungsi Delete
	return ret.Error(0)
}
