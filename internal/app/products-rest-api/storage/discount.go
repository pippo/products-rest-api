package storage

import (
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

type MySQLDiscountStorage struct{}

func NewMySQLDiscountStorage() *MySQLDiscountStorage {
	return &MySQLDiscountStorage{}
}

func (s *MySQLDiscountStorage) Add(discount models.Discount) error {
	// NOT implemented intentionally -- only used in tests
	return nil
}

func (s *MySQLDiscountStorage) LoadAll() ([]models.Discount, error) {
	db, err := dbConn()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	defer db.Close()

	rs, err := db.Query("SELECT sku, category, percent FROM discounts")
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
