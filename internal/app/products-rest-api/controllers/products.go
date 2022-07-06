package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/services"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
	"github.com/sirupsen/logrus"
)

type ProductListResponse struct {
	Products []models.DiscountedProduct `json:"products"`
}

func HandleListProducts(c *gin.Context) {

	// TODO: move to state
	productStorage := storage.NewMySQLProductStorage()
	discountStorage := storage.NewMySQLDiscountStorage()
	discountService := services.NewDiscountService(discountStorage)
	svc := services.NewProductService(productStorage, discountService)

	category := models.Category("boots") // TODO: read from request

	products, err := svc.ListProducts(services.ProductFilterCriteria{Category: category})
	if err != nil {
		logrus.WithError(err).Error("failed to list products")
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unknown error"})
		return
	}

	c.JSON(http.StatusOK, ProductListResponse{Products: products})
}
