package services

import (
	"fmt"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
)

type DiscountService struct {
	storage    storage.DiscountStorage
	bySKU      map[models.SKU]models.DiscountValue
	byCategory map[models.Category]models.DiscountValue
}

func NewDiscountService(storage storage.DiscountStorage) *DiscountService {
	return &DiscountService{storage: storage}
}

func (s *DiscountService) init() error {
	discounts, err := s.storage.LoadAll()
	if err != nil {
		return fmt.Errorf("failed to load discounts: %w", err)
	}

	s.bySKU = make(map[models.SKU]models.DiscountValue)
	s.byCategory = make(map[models.Category]models.DiscountValue)

	for _, d := range discounts {
		if d.Category != "" {
			s.byCategory[d.Category] = d.Value
		}
		if d.SKU != "" {
			s.bySKU[d.SKU] = d.Value
		}
	}

	return nil
}

func (s *DiscountService) LookupDiscount(p models.Product) *models.DiscountValue {
	if s.bySKU == nil {
		s.init()
	}

	v, ok := s.byCategory[p.Category]
	if ok {
		return &v
	}

	v, ok = s.bySKU[p.SKU]
	if ok {
		return &v
	}
	return nil
}
