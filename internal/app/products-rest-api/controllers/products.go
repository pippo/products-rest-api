package controllers

import (
	"net/http"
	"strconv"

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

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewProductsController(ps *services.ProductService) *ProductsController {
	return &ProductsController{ProductsService: ps}
}

func (p *ProductsController) HandleListProducts(c *gin.Context) {
	categoryStr := c.Query("category")
	maxPriceStr := c.DefaultQuery("priceLessThan", "0")

	maxPriceInt, err := strconv.ParseInt(maxPriceStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid priceLessThan provided"})
		return
	}

	products, err := p.ProductsService.ListProducts(services.ProductFilterCriteria{
		Category: models.Category(categoryStr),
		MaxPrice: models.Price(maxPriceInt),
	})
	if err != nil {
		logrus.WithError(err).Error("failed to list products")
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Unknown error"})
		return
	}

	c.JSON(http.StatusOK, ProductListResponse{Products: products})
}
