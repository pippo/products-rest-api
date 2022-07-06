package storage

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
)

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

type MySQLProductStorage struct{}

func NewMySQLProductStorage() *MySQLProductStorage {
	return &MySQLProductStorage{}
}

func (s *MySQLProductStorage) Add(product models.Product) error {
	// NOT implemented intentionally -- only used in tests
	return nil
}

func (s *MySQLProductStorage) ListProducts(category models.Category, maxPrice models.Price, limit int) ([]models.Product, error) {
	db, err := dbConn()
	if err != nil {
		return nil, fmt.Errorf("failed to init DB connection: %w", err)
	}
	defer db.Close()

	q := sq.Select("sku, category, pname, price").From("products").Limit(uint64(limit))
	if category != "" {
		q = q.Where(sq.Eq{"category": category})
	}
	if maxPrice != 0 {
		q = q.Where(sq.Lt{"price": maxPrice})
	}

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate SQL query: %w", err)
	}

	rs, err := db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to run SQL query: %w", err)
	}

	result := []models.Product{}

	for rs.Next() {
		var p models.Product
		if err = rs.Scan(&p.SKU, &p.Category, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}
