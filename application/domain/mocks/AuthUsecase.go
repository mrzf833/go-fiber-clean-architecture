package mocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/auth/request"
)

type AuthUsecase struct {
	mock.Mock
}

func (m *AuthUsecase) Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error) {
	ret := m.Called(c.Context(), request.Username)
	var r0 map[string]interface{}
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]interface{})
	}
	return r0, ret.Error(1)
}