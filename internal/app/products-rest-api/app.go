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

func (a *ProductsRestAPIApp) Configure() error {
	db, err := storage.ConnectToMySQL()
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
