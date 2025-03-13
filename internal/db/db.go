package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDB initializes a PostgreSQL connection
func NewPostgresDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
