package inventory

import "gorm.io/gorm"

// Repository handles database operations for inventory
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new inventory Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
