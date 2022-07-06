package app

import (
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

func (a *ProductsRestAPIApp) Configure() {
	productStorage := storage.NewMySQLProductStorage()
	discountStorage := storage.NewMySQLDiscountStorage()
	discountService := services.NewDiscountService(discountStorage)

	a.ProductService = services.NewProductService(productStorage, discountService)
}

func (a *ProductsRestAPIApp) Run() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	ctrl := controllers.NewProductsController(a.ProductService)
	router.GET("/products", ctrl.HandleListProducts)

	router.Run()
}
