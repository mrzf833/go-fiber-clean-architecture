package domain

import (
	"github.com/gofiber/fiber/v2"
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
	GetAll(ctx *fiber.Ctx) ([]FileSave, error)
	Create(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Update(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Delete(ctx *fiber.Ctx, id int64) (error)
}

type FileSaveRepository interface {
	GetAll(ctx *fiber.Ctx) ([]FileSave, error)
	Create(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Update(ctx *fiber.Ctx, fileSave FileSave) (FileSave, error)
	Delete(ctx *fiber.Ctx, id int64) (error)
}