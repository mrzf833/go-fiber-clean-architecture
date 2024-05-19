package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
	"time"
)

type AuthRepository struct {
	mock.Mock
}

func (m *AuthRepository) CreateToken(ctx context.Context, auth domain.Auth, exp time.Duration) (domain.Auth, error) {
	// pemanggilan fungsi Create dengan parameter ctx dan category menggunakan mock
	ret := m.Called(ctx, auth, exp)

	// pengembalian value dari fungsi Create
	var r0 domain.Auth

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.Auth)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}

func (m *AuthRepository) DeleteToken(ctx context.Context, username string) error {
	// pemanggilan fungsi GetByID dengan parameter ctx dan username menggunakan mock
	ret := m.Called(ctx, username)

	// pengembalian value error
	return ret.Error(0)
}

func (m *AuthRepository) GetToken(ctx context.Context, username string) (domain.Auth, error) {
	// pemanggilan fungsi GetByID dengan parameter ctx dan username menggunakan mock
	ret := m.Called(ctx, username)

	// pengembalian value dari fungsi GetByID
	var r0 domain.Auth

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(1) == nil {
		r0 = ret.Get(0).(domain.Auth)
	}

	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}