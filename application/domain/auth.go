package domain

import (
	"context"
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
	Login(ctx context.Context, request request.AuthCreateRequest) (map[string]interface{}, error)
	//Logout(ctx context.Context) error
	//User(ctx context.Context, category Auth) (Auth, error)
}

type AuthRepository interface {
	GetByUsername(ctx context.Context, username string) (Auth, error)
}