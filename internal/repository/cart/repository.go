package cart

import "gorm.io/gorm"

// Repository handles database operations for shopping carts
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new cart Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
