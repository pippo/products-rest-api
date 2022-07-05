package storage

import "github.com/pippo/products-rest-api/internal/app/products-rest-api/models"

type DiscountStorage interface {
	Add(models.Discount) error
	LoadAll() ([]models.Discount, error)
}

type InMemoryDiscountStorage struct {
	storage []models.Discount
}

func NewInMemoryDiscountStorage() *InMemoryDiscountStorage {
	return &InMemoryDiscountStorage{storage: []models.Discount{}}
}

func (s *InMemoryDiscountStorage) Add(discount models.Discount) error {
	s.storage = append(s.storage, discount)
	return nil
}

func (s *InMemoryDiscountStorage) LoadAll() ([]models.Discount, error) {
	return s.storage, nil
}
