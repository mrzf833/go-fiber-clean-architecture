package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/file_save/delivery/http"
	gorm_mysql "go-fiber-clean-architecture/application/app/file_save/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/app/file_save/usecase"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/middleware"
)

func fileSaveRouterApi(api fiber.Router, validate *validator.Validate)  {
	// setup category repository
	fileSaveRepository := gorm_mysql.NewMysqlFileSaveRepository(config.DB)
	// setup category usecase
	categoryUseCase := usecase.NewFileSaveUseCase(fileSaveRepository)
	// setup category handler
	handler := http.NewFileSaveHandler(categoryUseCase, validate)


	// setup routes
	// with middleware auth jwt
	fileSaveApi := api.Group("/file-save", middleware.AuthMiddleware()...)
	fileSaveApi.Get("/:id", handler.GetByID)
	fileSaveApi.Get("/", handler.GetAll)
	fileSaveApi.Post("/", handler.Create)
	fileSaveApi.Post("/:id", handler.Update)
	fileSaveApi.Delete("/:id", handler.Delete)
}
