package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// CategoryRepository handles database operations for categories
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new CategoryRepository
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll retrieves all categories
func (r *CategoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Order("name").Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	return categories, nil
}

// GetByID retrieves a category by its ID
func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}
	return &category, nil
}

// Create inserts a new category
func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}

// Delete removes a category by ID
func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Category{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
