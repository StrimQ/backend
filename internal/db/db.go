package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func pgConnectionString(host string, port int, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
}

// NewPostgresDB initializes a PostgreSQL connection
func NewPostgresDB(host string, port int, user, password, dbName string) (*sql.DB, error) {
	pgConnectString := pgConnectionString(host, port, user, password, dbName)
	sqlDB, err := sql.Open("postgres", pgConnectString)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return sqlDB, nil
}
