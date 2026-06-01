package storesetting

import "gorm.io/gorm"

// Repository handles database operations for store settings
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new store setting Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
