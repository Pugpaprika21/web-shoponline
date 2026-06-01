package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all menus ordered by sort_order
func (r *Repository) GetAll(ctx context.Context) ([]model.Menu, error) {
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
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Menu, int64, error) {
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
func (r *Repository) GetTopLevel(ctx context.Context) ([]model.Menu, error) {
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
func (r *Repository) GetChildren(ctx context.Context, parentID string) ([]model.Menu, error) {
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
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Menu, error) {
	var m model.Menu
	err := r.db.WithContext(ctx).
		Preload("RoleIDs").
		Where("id = ?", id).
		First(&m).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get menu by id: %w", err)
	}
	return &m, nil
}
