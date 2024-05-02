package gorm_mysql

import (
	"context"
	"errors"
	"go-fiber-clean-architecture/application/domain"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	Db *gorm.DB
}

func NewMysqlAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &mysqlAuthRepository{db}
}

func (r *mysqlAuthRepository) GetByUsername(ctx context.Context, username string) (domain.Auth, error) {
	var auth domain.Auth
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).First(&auth, "username = ?", username).Error

	// jika data tidak ditemukan maka akan mengembalikan error not found
	if auth.ID == 0 {
		return auth, errors.New("username or password is wrong")
	}
	return auth, err
}