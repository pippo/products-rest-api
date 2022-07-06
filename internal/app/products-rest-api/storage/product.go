package storage

import (
	"database/sql"
	"fmt"

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
	return nil
}

func (s *MySQLProductStorage) ListProducts(category models.Category, maxPrice models.Price, limit int) ([]models.Product, error) {
	db, err := dbConn()
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer db.Close()

	rs, err := db.Query("SELECT sku, category, pname, price FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
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

func dbConn() (*sql.DB, error) {
	// TODO: read credentials from ENV or Vault
	db, err := sql.Open("mysql", "root:@tcp(db:3306)/products_rest_api")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}
