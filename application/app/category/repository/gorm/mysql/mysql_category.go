package gorm_mysql

import (
	"context"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	Db *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &mysqlCategoryRepository{db}
}

func (r *mysqlCategoryRepository) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	var category domain.Category
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).First(&category, id).Error

	// jika data tidak ditemukan maka akan mengembalikan error not found
	if category.ID == 0 {
		return category, exception.ErrNotFound
	}
	return category, err
}

func (r *mysqlCategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).Find(&categories).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return categories, err
}

func (r *mysqlCategoryRepository) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	// insert data ke database menggunakan gorm
	err := r.Db.WithContext(ctx).Create(&category).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return category, err
}

func (r *mysqlCategoryRepository) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	// update data ke database menggunakan gorm
	err := r.Db.WithContext(ctx).Updates(&category).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return category, err
}

func (r *mysqlCategoryRepository) Delete(ctx context.Context, id int64) error {
	// delete data ke database menggunakan gorm
	err := r.Db.WithContext(ctx).Delete(&domain.Category{}, id).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return err
}