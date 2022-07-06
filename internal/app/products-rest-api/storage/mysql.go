package storage

import (
	"database/sql"
	"fmt"
)

func ConnectToMySQL(host, port, user, pass string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/products_rest_api", user, pass, host, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}
