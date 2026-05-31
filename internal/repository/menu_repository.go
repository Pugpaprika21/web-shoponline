package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// MenuRepository handles database operations for menus
type MenuRepository struct {
	db *gorm.DB
}

// NewMenuRepository creates a new MenuRepository
func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

// GetAll retrieves all menus ordered by sort_order
func (r *MenuRepository) GetAll(ctx context.Context) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.WithContext(ctx).
		Preload("RoleIDs").
		Order("sort_order").
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query menus: %w", err)
	}
	return menus, nil
}

// GetAllPaginated retrieves menus with pagination and optional search
func (r *MenuRepository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Menu, int64, error) {
	var menus []model.Menu
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Menu{})
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("RoleIDs")
	if search != "" {
		q = q.Where("name ILIKE ?", "%"+search+"%")
	}
	err := q.Order("sort_order").
		Offset(offset).Limit(limit).
		Find(&menus).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query menus: %w", err)
	}
	return menus, total, nil
}

// GetTopLevel retrieves only top-level menus (no parent)
func (r *MenuRepository) GetTopLevel(ctx context.Context) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.WithContext(ctx).
		Where("parent_id IS NULL").
		Order("sort_order").
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query top-level menus: %w", err)
	}
	return menus, nil
}

// GetChildren retrieves child menus for a parent
func (r *MenuRepository) GetChildren(ctx context.Context, parentID string) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("sort_order").
		Find(&menus).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query child menus: %w", err)
	}
	return menus, nil
}

// GetByID retrieves a menu by ID
func (r *MenuRepository) GetByID(ctx context.Context, id string) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).
		Preload("RoleIDs").
		Where("id = ?", id).
		First(&menu).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get menu by id: %w", err)
	}
	return &menu, nil
}

// Create inserts a new menu
func (r *MenuRepository) Create(ctx context.Context, menu *model.Menu) error {
	if err := r.db.WithContext(ctx).Create(menu).Error; err != nil {
		return fmt.Errorf("failed to create menu: %w", err)
	}
	return nil
}

// Update modifies an existing menu
func (r *MenuRepository) Update(ctx context.Context, menu *model.Menu) error {
	if err := r.db.WithContext(ctx).Save(menu).Error; err != nil {
		return fmt.Errorf("failed to update menu: %w", err)
	}
	return nil
}

// Delete removes a menu by ID
func (r *MenuRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Menu{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete menu: %w", err)
	}
	return nil
}
