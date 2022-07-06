package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/services"
	"github.com/sirupsen/logrus"
)

type ProductsController struct {
	ProductsService *services.ProductService
}

type ProductListResponse struct {
	Products []models.DiscountedProduct `json:"products"`
}

func NewProductsController(ps *services.ProductService) *ProductsController {
	return &ProductsController{ProductsService: ps}
}

func (p *ProductsController) HandleListProducts(c *gin.Context) {
	category := models.Category("boots") // TODO: read from request

	products, err := p.ProductsService.ListProducts(services.ProductFilterCriteria{Category: category})
	if err != nil {
		logrus.WithError(err).Error("failed to list products")
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unknown error"})
		return
	}

	c.JSON(http.StatusOK, ProductListResponse{Products: products})
}
