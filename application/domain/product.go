package domain

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Product struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement" mapstructure:"id"`
	CategoryId int64 `json:"category_id" gorm:"column:category_id" mapstructure:"category_id"`
	Name string `json:"name" gorm:"column:name" validate:"required" mapstructure:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime" mapstructure:"updated_at"`
}

//type Product struct {
//	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
//	CategoryId int64 `json:"category_id" gorm:"column:category_id"`
//	Name string `json:"name" gorm:"column:name" validate:"required"`
//	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime"`
//}

func (c *Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	GetByID(ctx context.Context, id int64) (Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int64) error
}

type ProductUseCase interface {
	GetByID(ctx context.Context, id int64) (Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int64) error
}

type ProductHandler interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
