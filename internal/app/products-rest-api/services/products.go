package services

import (
	"fmt"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
)

const ListMaxItems = 5

type ProductService struct {
	storage storage.ProductStorage
}

type ProductFilterCriteria struct {
	Category models.Category
	MaxPrice models.Price
}

func NewProductService(storage storage.ProductStorage) *ProductService {
	return &ProductService{storage: storage}
}

func (s *ProductService) ListProducts(filter ProductFilterCriteria) ([]models.DiscountedProduct, error) {
	result := []models.DiscountedProduct{}

	products, err := s.storage.ListProducts(filter.Category, filter.MaxPrice, ListMaxItems)
	if err != nil {
		return result, fmt.Errorf("failed to list products: %w", err)
	}

	// TODO: apply discounts
	for _, p := range products {
		// TODO: add helper: func applyDiscount(Product, Discount) DiscountedProduct
		result = append(result, models.DiscountedProduct{
			SKU:      p.SKU,
			Category: p.Category,
			Name:     p.Name,
			Price: models.PriceWithDiscount{
				Original:           p.Price,
				Final:              p.Price,
				DiscountPercentage: "0%",
				Currency:           models.CurrencyEUR,
			},
		})
	}

	return result, nil
}
