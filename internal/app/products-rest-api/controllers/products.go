package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/services"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
)

func HandleListProducts(c *gin.Context) {

	// TODO: move to state
	storage := storage.NewInMemoryProductStorage()
	svc := services.NewProductService(storage)

	category := models.Category("boots") // TODO: read from request

	products, err := svc.ListProducts(services.ProductFilterCriteria{Category: category})
	if err != nil {
		// TODO: proper handling
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	_ = products
}
