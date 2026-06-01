package product

import "gorm.io/gorm"

// Repository handles database operations for products
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new product Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
