package gorm_mysql

import (
	"context"
	"fmt"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db}
}

func (r mysqlProductRepository) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	var product domain.Product
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).First(&product, id).Error

	// jika data tidak ditemukan maka akan mengembalikan error not found
	if product.ID == 0 {
		return product, exception.ErrNotFound
	}
	return product, err
}

func (r mysqlProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).Find(&products).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return products, err
}

func (r mysqlProductRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	// insert data ke database menggunakan gorm
	fmt.Println(product)
	err := r.Db.WithContext(ctx).Create(&product).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return product, err
}

func (r mysqlProductRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	// update data ke database menggunakan gorm
	err := r.Db.WithContext(ctx).Updates(&product).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return product, err
}

func (r mysqlProductRepository) Delete(ctx context.Context, id int64) error {
	// delete data ke database menggunakan gorm
	err := r.Db.WithContext(ctx).Delete(&domain.Product{}, id).Error
	// tanpa pengecekan karena jika data tidak ditemukan maka akan mengembalikan array kosong
	return err
}