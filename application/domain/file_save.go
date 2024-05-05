package domain

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type FileSave struct {
	ID int64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoCreateTime"`
}

func (c *FileSave) TableName() string {
	return "file_saves"
}

type FileSaveUsecase interface {
	GetByID(ctx *fiber.Ctx, id int64) (FileSave, error)
	GetAll(ctx *fiber.Ctx) ([]FileSave, error)
	Create(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Update(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Delete(ctx *fiber.Ctx, id int64) (error)
}

type FileSaveRepository interface {
	GetAll(ctx *fiber.Ctx) ([]FileSave, error)
	GetByID(ctx *fiber.Ctx, id int64) (FileSave, error)
	Create(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Update(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Delete(ctx *fiber.Ctx, id int64) (error)
	GetDb() (db *gorm.DB)
}

type FileSaveHandler interface {
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}