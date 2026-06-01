package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all roles with permissions
func (r *Repository) GetAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Order("name").
		Find(&roles).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query roles: %w", err)
	}
	return roles, nil
}

// GetAllPaginated retrieves roles with pagination and optional search
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Role{})
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Permissions")
	if search != "" {
		q = q.Where("name ILIKE ?", "%"+search+"%")
	}
	err := q.Order("name").
		Offset(offset).Limit(limit).
		Find(&roles).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query roles: %w", err)
	}
	return roles, total, nil
}

// GetByID retrieves a role by ID with permissions
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Role, error) {
	var rl model.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Where("id = ?", id).
		First(&rl).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get role by id: %w", err)
	}
	return &rl, nil
}

// GetAllPermissions retrieves all available permissions
func (r *Repository) GetAllPermissions(ctx context.Context) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.WithContext(ctx).Order("module, name").Find(&permissions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query permissions: %w", err)
	}
	return permissions, nil
}
