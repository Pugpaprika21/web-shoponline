package menu

import "gorm.io/gorm"

// Repository handles database operations for menus
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new menu Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
