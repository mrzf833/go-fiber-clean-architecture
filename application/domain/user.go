package domain

import (
	"context"
	"time"
)

type User struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime"`
}

func (c *User) TableName() string {
	return "users"
}


type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (User, error)
}