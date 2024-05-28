package domain

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"io"
	"time"
)

type Category struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime"`
}

func (c *Category) TableName() string {
	return "categories"
}

type CategoryUseCase interface {
	GetByID(ctx context.Context, id int64) (Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category Category) (Category, error)
	Update(ctx context.Context, category Category) (Category, error)
	Delete(ctx context.Context, id int64) error
	CreateWithCsv(ctx context.Context, file io.Reader, idTrackerCategory int64)
	CreateWithCsvQueue(ctx context.Context, file io.Reader) error
}

type CategoryRepository interface {
	GetByID(ctx context.Context, id int64) (Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	Create(ctx context.Context, category Category) (Category, error)
	Update(ctx context.Context, category Category) (Category, error)
	Delete(ctx context.Context, id int64) (error)
	CreateAll(ctx context.Context, category []Category) ([]Category, error)
	CreateInBatches(ctx context.Context, category []Category, size int) error
}

type CategoryHandler interface {
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	CreateWithCsv(c *fiber.Ctx) error
	CreateWithCsvQueue(c *fiber.Ctx) error
}

type QueueCategoryJob interface {
	HandleCategoryCreateWithCsvQueue(ctx context.Context, t *asynq.Task) error
}