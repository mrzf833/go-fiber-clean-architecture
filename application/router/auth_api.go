package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/auth/delivery/http"
	gorm_mysql "go-fiber-clean-architecture/application/app/auth/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/app/auth/usecase"
	"go-fiber-clean-architecture/application/config"
)

func authRouterApi(api fiber.Router, validate *validator.Validate)  {
	// setup category repository
	authRepository := gorm_mysql.NewMysqlAuthRepository(config.DB)
	// setup category usecase
	authUseCase := usecase.NewAuthUseCase(authRepository)
	// setup category router
	http.NewAuthHandler(api, authUseCase,validate)
}