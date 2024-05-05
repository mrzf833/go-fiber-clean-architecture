package mocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"go-fiber-clean-architecture/application/app/auth/request"
)

type AuthUsecase struct {
	mock.Mock
}

func (m *AuthUsecase) Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error) {
	ret := m.Called(c, request)
	var r0 map[string]interface{}
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]interface{})
	}
	return r0, ret.Error(1)
}

func (m *AuthUsecase) User(c *fiber.Ctx) jwt.MapClaims {
	ret := m.Called(c)
	var r0 jwt.MapClaims
	r0 = ret.Get(0).(jwt.MapClaims)
	return r0
}

func (m *AuthUsecase) Logout(c *fiber.Ctx) {
	m.Called(c)
}