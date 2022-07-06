package services

import (
	"fmt"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
)

const ListMaxItems = 5

type ProductService struct {
	storage         storage.ProductStorage
	discountService *DiscountService
}

type ProductFilterCriteria struct {
	Category models.Category
	MaxPrice models.Price
}

func NewProductService(storage storage.ProductStorage, ds *DiscountService) *ProductService {
	return &ProductService{
		storage:         storage,
		discountService: ds,
	}
}

func (s *ProductService) ListProducts(filter ProductFilterCriteria) ([]models.DiscountedProduct, error) {
	result := []models.DiscountedProduct{}

	products, err := s.storage.ListProducts(filter.Category, filter.MaxPrice, ListMaxItems)
	if err != nil {
		return result, fmt.Errorf("failed to list products: %w", err)
	}

	for _, p := range products {
		dv := s.discountService.LookupDiscount(p)

		dp, err := p.ApplyDiscount(dv)
		if err != nil {
			return nil, fmt.Errorf("failed to apply discount: %w", err)
		}

		result = append(result, *dp)
	}

	return result, nil
}
