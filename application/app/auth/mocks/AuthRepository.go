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

func (m *AuthRepository) CreateToken(ctx context.Context, username string, auth domain.Auth, exp time.Duration) error {
	// pemanggilan fungsi GetByID dengan parameter ctx dan username menggunakan mock
	ret := m.Called(ctx, username, auth, exp)

	// pengembalian value error
	return ret.Error(0)
}

func (m *AuthRepository) DeleteToken(ctx context.Context, username string) error {
	// pemanggilan fungsi GetByID dengan parameter ctx dan username menggunakan mock
	ret := m.Called(ctx, username)

	// pengembalian value error
	return ret.Error(0)
}