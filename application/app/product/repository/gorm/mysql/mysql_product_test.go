package gorm_mysql_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	gorm_mysql "go-fiber-clean-architecture/application/app/product/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	db, mock := utils.NewMockDB()

	mockProduct := []domain.Product{
		{
			ID:   1,
			Name: "Product 1",
			CategoryId: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:   2,
			Name: "Product 2",
			CategoryId: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "category_id", "created_at", "updated_at"}).
		AddRow(mockProduct[0].ID, mockProduct[0].Name, mockProduct[0].CategoryId, mockProduct[0].CreatedAt, mockProduct[0].UpdatedAt).
		AddRow(mockProduct[1].ID, mockProduct[1].Name, mockProduct[0].CategoryId, mockProduct[1].CreatedAt, mockProduct[1].UpdatedAt)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlProductRepository(db)
	products, err := a.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, products, 2)
}

func TestGetById_success(t *testing.T) {
	db, mock := utils.NewMockDB()

	mockProduct := domain.Product{
		ID:   1,
		Name: "product 1",
		CategoryId: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "category_id", "created_at", "updated_at"}).
		AddRow(mockProduct.ID, mockProduct.Name, mockProduct.CategoryId, mockProduct.CreatedAt, mockProduct.UpdatedAt)

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlProductRepository(db)

	category, err := a.GetByID(context.Background(), mockProduct.ID)

	assert.NoError(t, err)
	assert.Equal(t, category.ID, mockProduct.ID)
}

func TestGetById_notFound(t *testing.T) {
	db, mock := utils.NewMockDB()

	mockProduct := domain.Product{
		ID:   1,
		Name: "Product 1",
		CategoryId: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "category_id", "created_at", "updated_at"})

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlProductRepository(db)

	_, err := a.GetByID(context.Background(), mockProduct.ID)

	assert.Error(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreate_success(t *testing.T) {
	db, mock := utils.NewMockDB()

	now := time.Now()
	prod := domain.Product{
		Name: "product 1",
		CategoryId: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT INTO `products`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlProductRepository(db)

	product, err := a.Create(context.TODO(), prod)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), product.ID)
}

func TestUpdate_success(t *testing.T) {
	db, mock := utils.NewMockDB()

	now := time.Now()
	prod := domain.Product{
		ID: 1,
		Name: "product 1",
		CategoryId: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "UPDATE `products`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlProductRepository(db)

	product, err := a.Update(context.TODO(), prod)
	assert.NoError(t, err)
	assert.Equal(t, prod.ID, product.ID)
}

func TestDelete_success(t *testing.T) {
	db, mock := utils.NewMockDB()

	now := time.Now()
	prod := domain.Product{
		ID: 1,
		Name: "product 1",
		CategoryId: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "DELETE FROM `products`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlProductRepository(db)

	err := a.Delete(context.TODO(), prod.ID)
	assert.NoError(t, err)
}
