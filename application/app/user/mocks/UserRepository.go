package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/domain"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	// pemanggilan fungsi GetByID dengan parameter ctx dan id menggunakan mock
	ret := m.Called(ctx, username)

	// pengembalian value dari fungsi GetByID
	var r0 domain.User

	// jika value yang dikembalikan tidak nil maka value tersebut di assign ke r0
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domain.User)
	}
	// pengembalian value r0 dan error
	return r0, ret.Error(1)
}
