package domain

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app/auth/request"
	"time"
)

type Auth struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Token string `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthUseCase interface {
	Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error)
	User(c *fiber.Ctx) jwt.MapClaims
	Logout(c *fiber.Ctx)
}

type AuthRepository interface {
	CreateToken(ctx context.Context, data Auth, exp time.Duration) (Auth, error)
	DeleteToken(ctx context.Context, username string) error
	GetToken(ctx context.Context, username string) (Auth, error)
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	User(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}