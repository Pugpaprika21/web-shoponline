package order

import "gorm.io/gorm"

// Repository handles database operations for orders
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new order Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
