package gorm_mysql_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	gorm_mysql "go-fiber-clean-architecture/application/app/file_save/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	mockFileSaves := []domain.FileSave{
		{
			ID:   1,
			Name: "file 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:   2,
			Name: "file 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(mockFileSaves[0].ID, mockFileSaves[0].Name, mockFileSaves[0].CreatedAt, mockFileSaves[0].UpdatedAt).
		AddRow(mockFileSaves[1].ID, mockFileSaves[1].Name, mockFileSaves[1].CreatedAt, mockFileSaves[1].UpdatedAt)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var categories []domain.FileSave
	var err error
	app.Use(func(c *fiber.Ctx) error {
		categories, err = a.GetAll(c)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestGetById_success(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	mockFileSave := domain.FileSave{
		ID:   1,
		Name: "File 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(mockFileSave.ID, mockFileSave.Name, mockFileSave.CreatedAt, mockFileSave.UpdatedAt)

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var fileSave domain.FileSave
	var err error
	app.Use(func(c *fiber.Ctx) error {
		fileSave, err = a.GetByID(c, mockFileSave.ID)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fileSave.ID, mockFileSave.ID)
}

func TestGetById_notFound(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	mockFileSave := domain.FileSave{
		ID:   1,
		Name: "File 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})

	mock.ExpectQuery("WHERE").WillReturnRows(rows)

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var fileSave domain.FileSave
	var err error
	app.Use(func(c *fiber.Ctx) error {
		fileSave, err = a.GetByID(c, mockFileSave.ID)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.Error(t, err)
	assert.Equal(t, fileSave.ID, int64(0))
}

func TestCreate_success(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	now := time.Now()
	fileSave := domain.FileSave{
		Name: "file 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "INSERT INTO `file_saves`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var err error
	app.Use(func(c *fiber.Ctx) error {
		fileSave, err = a.Create(c, fileSave)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), fileSave.ID)
}

func TestUpdate_success(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	now := time.Now()
	fileSaveBefore := domain.FileSave{
		ID: 1,
		Name: "category 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "UPDATE `file_saves`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var fileSaveAfter domain.FileSave
	var err error
	app.Use(func(c *fiber.Ctx) error {
		fileSaveAfter, err = a.Update(c, fileSaveBefore)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fileSaveBefore.ID, fileSaveAfter.ID)
}

func TestDelete_success(t *testing.T) {
	db, mock := utils.NewMockDB()
	app := fiber.New()

	now := time.Now()
	fileSaveBefore := domain.FileSave{
		ID: 1,
		Name: "file 1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := "DELETE FROM `file_saves`"
	mock.ExpectBegin()
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	a := gorm_mysql.NewMysqlFileSaveRepository(db)

	var err error
	app.Use(func(c *fiber.Ctx) error {
		err = a.Delete(c, fileSaveBefore.ID)
		return nil
	})
	req := httptest.NewRequest(fiber.MethodGet, "/", nil)
	app.Test(req)
	assert.NoError(t, err)
}
