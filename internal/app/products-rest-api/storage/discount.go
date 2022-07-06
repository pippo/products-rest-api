package storage

import (
	"database/sql"
	"fmt"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
)

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

type MySQLDiscountStorage struct {
	db *sql.DB
}

func NewMySQLDiscountStorage(db *sql.DB) *MySQLDiscountStorage {
	return &MySQLDiscountStorage{db: db}
}

func (s *MySQLDiscountStorage) Add(discount models.Discount) error {
	// NOT implemented intentionally -- only used in tests
	return nil
}

func (s *MySQLDiscountStorage) LoadAll() ([]models.Discount, error) {
	rs, err := s.db.Query("SELECT sku, category, percent FROM discounts")
	if err != nil {
		return nil, fmt.Errorf("failed to query discounts: %w", err)
	}

	result := []models.Discount{}

	for rs.Next() {
		var d models.Discount
		var sku *models.SKU
		var cat *models.Category
		if err = rs.Scan(&sku, &cat, &d.Value); err != nil {
			return nil, err
		}
		if sku != nil {
			d.SKU = *sku
		}
		if cat != nil {
			d.Category = *cat
		}
		result = append(result, d)
	}

	return result, nil
}
