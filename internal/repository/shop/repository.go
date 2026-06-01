package shop

import "gorm.io/gorm"

// Repository handles database operations for shops
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new shop Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
