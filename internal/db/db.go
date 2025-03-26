package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func pgConnectionString(host string, port int, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
}

// NewPostgresDB initializes a PostgreSQL connection
func NewPostgresDB(ctx context.Context, host string, port int, user, password, dbName string) (*pgxpool.Pool, error) {
	pgConnectString := pgConnectionString(host, port, user, password, dbName)
	return pgxpool.New(ctx, pgConnectString)
}
