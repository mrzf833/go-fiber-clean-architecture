package gorm_mysql_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	gorm_mysql "go-fiber-clean-architecture/application/app/category/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	db, mock := helper.NewMockDB()

	mockCategories := []domain.Category{
		{
			ID:   1,
			Name: "Category 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:   2,
			Name: "Category 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(mockCategories[0].ID, mockCategories[0].Name, mockCategories[0].CreatedAt, mockCategories[0].UpdatedAt).
		AddRow(mockCategories[1].ID, mockCategories[1].Name, mockCategories[1].CreatedAt, mockCategories[1].UpdatedAt)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlCategoryRepository(db)
	categories, err := a.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestGetById_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	mockCategory := domain.Category{
		ID:   1,
		Name: "Category 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(mockCategory.ID, mockCategory.Name, mockCategory.CreatedAt, mockCategory.UpdatedAt)

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	category, err := a.GetByID(context.Background(), mockCategory.ID)

	assert.NoError(t, err)
	assert.Equal(t, category.ID, mockCategory.ID)
}

func TestGetById_notFound(t *testing.T) {
	db, mock := helper.NewMockDB()

	mockCategory := domain.Category{
		ID:   1,
		Name: "Category 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	_, err := a.GetByID(context.Background(), mockCategory.ID)

	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreate_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	now := time.Now()
	cat := domain.Category{
		Name: "category 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT INTO `categories`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	category, err := a.Create(context.TODO(), cat)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), category.ID)
}


func TestUpdate_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	now := time.Now()
	cat := domain.Category{
		ID: 1,
		Name: "category 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "UPDATE `categories`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	category, err := a.Update(context.TODO(), cat)
	assert.NoError(t, err)
	assert.Equal(t, cat.ID, category.ID)
}


func TestDelete_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	now := time.Now()
	cat := domain.Category{
		ID: 1,
		Name: "category 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "DELETE FROM `categories`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	err := a.Delete(context.TODO(), cat.ID)
	assert.NoError(t, err)
}

func TestCreateAll_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	now := time.Now()
	cat := []domain.Category{
		{
			Name:      "category 1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:      "category 2",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	query := "INSERT INTO `categories`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	category, err := a.CreateAll(context.TODO(), cat)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), category[0].ID)
	assert.Equal(t, int64(13), category[1].ID)
}

func TestCreateInBatches_success(t *testing.T) {
	db, mock := helper.NewMockDB()

	now := time.Now()
	cat := []domain.Category{
		{
			Name:      "category 1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:      "category 2",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:      "category 3",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	query := "INSERT INTO `categories`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 2))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(14, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlCategoryRepository(db)

	err := a.CreateInBatches(context.TODO(), cat, 2)
	assert.NoError(t, err)
}