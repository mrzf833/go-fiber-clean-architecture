package mysql

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/domain"
	"gorm.io/gorm"
)

type mysqlFileSaveRepository struct {
	Db *gorm.DB
}

func NewMysqlFileSaveRepository(db *gorm.DB) domain.FileSaveRepository {
	return &mysqlFileSaveRepository{db}
}

func (r *mysqlFileSaveRepository) GetAll(c *fiber.Ctx) ([]domain.FileSave, error) {
	var fileSaves []domain.FileSave
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(c.Context()).Find(&fileSaves).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return fileSaves, err
}

func (r *mysqlFileSaveRepository) Create(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// insert data ke database menggunakan gorm
	err := r.Db.WithContext(c.Context()).Create(&fileSave).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return fileSave, err
}

func (r *mysqlFileSaveRepository) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// update data ke database menggunakan gorm
	err := r.Db.WithContext(c.Context()).Updates(&fileSave).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return fileSave, err
}

func (r *mysqlFileSaveRepository) Delete(c *fiber.Ctx, id int64) error {
	// delete data ke database menggunakan gorm
	err := r.Db.WithContext(c.Context()).Delete(&domain.FileSave{}, id).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return err
}
