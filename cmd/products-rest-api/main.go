package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/controllers"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.GET("/products", controllers.HandleListProducts)
	router.Run()
}
