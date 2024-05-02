package domain

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/auth/request"
	"time"
)

type Auth struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime"`
}

func (c *Auth) TableName() string {
	return "users"
}

type AuthUseCase interface {
	Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error)
}

type AuthRepository interface {
	GetByUsername(ctx context.Context, username string) (Auth, error)
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	User(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}