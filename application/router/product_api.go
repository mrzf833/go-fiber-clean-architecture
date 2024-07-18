package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/app/product/delivery/http"
	"go-fiber-clean-architecture/application/app/product/repository/elasticsearch"
	gorm_mysql "go-fiber-clean-architecture/application/app/product/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/app/product/usecase"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/middleware"
)

func productRouterApi(api fiber.Router, validate *validator.Validate)  {
	// setup product repository
	mysqlProductRepository := gorm_mysql.NewMysqlProductRepository(config.DB)
	elastichProductRepository := elasticsearch.NewElasticProductRepository(config.ElasticDb)
	// setup product usecase
	productUseCase := usecase.NewProductUseCase(mysqlProductRepository, elastichProductRepository)
	// setup product handler
	handler := http.NewProductHandler(productUseCase, validate)

	// setup routes
	// with middleware auth jwt
	productApi := api.Group("/product", middleware.AuthMiddleware()...)
	//productApi.Get("/", handler.GetAll)
	productApi.Get("/", handler.GetAll)
	productApi.Post("/", handler.Create)
	productApi.Get("/:id", handler.GetByID)
	productApi.Delete("/:id", handler.Delete)
}
