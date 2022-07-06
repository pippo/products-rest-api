package storage

import (
	"database/sql"
	"fmt"
)

func dbConn() (*sql.DB, error) {
	// TODO: read credentials from ENV or Vault
	db, err := sql.Open("mysql", "root:@tcp(db:3306)/products_rest_api")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}
