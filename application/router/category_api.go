package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/category/delivery/http"
	gorm_mysql "go-fiber-clean-architecture/application/app/category/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/app/category/usecase"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/middleware"
)

func categoryRouterApi(api fiber.Router, validate *validator.Validate)  {
	// setup category repository
	categoryRepository := gorm_mysql.NewMysqlCategoryRepository(config.DB)
	// setup category usecase
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	// setup category router
	// with middleware auth jwt
	categoryApi := api.Group("/category", middleware.AuthMiddleware()...)
	http.NewCategoryHandler(categoryApi, usecase.NewCategoryUseCase(categoryUseCase), validate)
}