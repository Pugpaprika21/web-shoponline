package category

import "gorm.io/gorm"

// Repository handles database operations for categories
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new category Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
