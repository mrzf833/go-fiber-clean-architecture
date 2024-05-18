package mysql

import (
	"context"
	"errors"
	"go-fiber-clean-architecture/application/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	// mengambil data dari database menggunakan gorm
	err := r.Db.WithContext(ctx).First(&user, "username = ?", username).Error

	// jika data tidak ditemukan maka akan mengembalikan error not found
	if user.ID == 0 {
		return user, errors.New("username not found")
	}
	return user, err
}