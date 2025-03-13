package repositories

import "gorm.io/gorm"

type SourceRepository struct {
	db *gorm.DB
}

func NewSourceRepository(db *gorm.DB) *SourceRepository {
	return &SourceRepository{db}
}
