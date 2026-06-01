package category

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all categories
func (r *Repository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Order("name").Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	return categories, nil
}

// GetByID retrieves a category by its ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Category, error) {
	var cat model.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&cat).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}
	return &cat, nil
}
