package role

import "gorm.io/gorm"

// Repository handles database operations for roles
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new role Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
