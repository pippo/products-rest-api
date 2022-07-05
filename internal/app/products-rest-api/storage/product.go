package storage

import "github.com/pippo/products-rest-api/internal/app/products-rest-api/models"

type ProductStorage interface {
	Add(product models.Product) error
	ListProducts(category models.Category, maxPrice models.Price, limit int) ([]models.Product, error)
}

type InMemoryProductStorage struct {
	storage map[models.SKU]models.Product
}

func NewInMemoryProductStorage() *InMemoryProductStorage {
	return &InMemoryProductStorage{storage: map[models.SKU]models.Product{}}
}

func (s *InMemoryProductStorage) Add(product models.Product) error {
	s.storage[product.SKU] = product
	return nil
}

func (s *InMemoryProductStorage) ListProducts(category models.Category, maxPrice models.Price, limit int) ([]models.Product, error) {
	result := []models.Product{}
	for _, item := range s.storage {
		if category != "" && category != item.Category {
			continue
		}
		if maxPrice > 0 && item.Price > maxPrice {
			continue
		}
		result = append(result, item)
		if len(result) >= limit {
			break
		}
	}
	return result, nil
}
