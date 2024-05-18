package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
	"testing"
	"time"
)

func TestGetByUsername_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	mockUser := domain.User{
		ID:   1,
		Username: "John",
		Password: "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).
		AddRow(mockUser.ID, mockUser.Username, mockUser.Password, mockUser.CreatedAt, mockUser.UpdatedAt)

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := NewMysqlUserRepository(db)

	auth, err := a.GetByUsername(context.Background(), mockUser.Username)

	assert.NoError(t, err)
	assert.Equal(t, auth.Username, mockUser.Username)
}

func TestGetByUsername_notFound(t *testing.T) {
	db, mock := helper.NewMockDB()

	mockUser := domain.User{
		ID:   1,
		Username: "John",
		Password: "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"})

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := NewMysqlUserRepository(db)

	_, err := a.GetByUsername(context.Background(), mockUser.Username)

	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}