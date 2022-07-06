package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/controllers"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/services"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
)

type ProductsRestAPIApp struct {
	ProductService *services.ProductService
}

func New() *ProductsRestAPIApp {
	return &ProductsRestAPIApp{}
}

func (a *ProductsRestAPIApp) Configure() error {
	mysqlHost := getEnvOrDefault("DB_HOST", "localhost")
	mysqlPort := getEnvOrDefault("DB_PORT", "3306")
	mysqlUser := getEnvOrDefault("DB_USER", "root")
	mysqlPass := getEnvOrDefault("DB_PASSWORD", "")

	db, err := storage.ConnectToMySQL(mysqlHost, mysqlPort, mysqlUser, mysqlPass)
	if err != nil {
		return err
	}

	productStorage := storage.NewMySQLProductStorage(db)
	discountStorage := storage.NewMySQLDiscountStorage(db)
	discountService := services.NewDiscountService(discountStorage)

	a.ProductService = services.NewProductService(productStorage, discountService)

	return nil
}

func (a *ProductsRestAPIApp) Run() {
	router := gin.Default()
	router.Use(gin.Recovery())

	ctrl := controllers.NewProductsController(a.ProductService)
	router.GET("/products", ctrl.HandleListProducts)

	router.Run()
}

func getEnvOrDefault(name, def string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return def
	}
	return val
}
