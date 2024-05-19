package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/auth/delivery/http"
	"go-fiber-clean-architecture/application/app/auth/repository/redis"
	"go-fiber-clean-architecture/application/app/auth/usecase"
	gorm_mysql "go-fiber-clean-architecture/application/app/user/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/middleware"
)

func authRouterApi(api fiber.Router, validate *validator.Validate)  {
	// setup category repository
	userRepository := gorm_mysql.NewMysqlUserRepository(config.DB)
	authRepository := redis.NewRedisAuthRepository(config.RedisDb)
	// setup category usecase
	authUseCase := usecase.NewAuthUseCase(authRepository, userRepository)
	// setup category router
	handler := http.NewAuthHandler(authUseCase,validate)

	// setup routes
	api.Post("/login", handler.Login)
	api.Get("/user", append(middleware.AuthMiddleware(), handler.User)...)
	api.Post("/logout", append(middleware.AuthMiddleware(), handler.Logout)...)
}