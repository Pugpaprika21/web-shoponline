package user

import "gorm.io/gorm"

// Repository handles database operations for users
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new user Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
