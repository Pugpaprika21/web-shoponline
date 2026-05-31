package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// RoleRepository handles database operations for roles
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a new RoleRepository
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// GetAll retrieves all roles with permissions
func (r *RoleRepository) GetAll(ctx context.Context) ([]model.Role, error) {
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
func (r *RoleRepository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Role, int64, error) {
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
func (r *RoleRepository) GetByID(ctx context.Context, id string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Where("id = ?", id).
		First(&role).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get role by id: %w", err)
	}
	return &role, nil
}

// Create inserts a new role
func (r *RoleRepository) Create(ctx context.Context, role *model.Role) error {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return fmt.Errorf("failed to create role: %w", err)
	}
	return nil
}

// Update modifies an existing role
func (r *RoleRepository) Update(ctx context.Context, role *model.Role) error {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}
	return nil
}

// UpdatePermissions replaces the permissions for a role
func (r *RoleRepository) UpdatePermissions(ctx context.Context, roleID string, permissions []model.Permission) error {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("id = ?", roleID).First(&role).Error; err != nil {
		return fmt.Errorf("failed to find role: %w", err)
	}
	if err := r.db.WithContext(ctx).Model(&role).Association("Permissions").Replace(permissions); err != nil {
		return fmt.Errorf("failed to update permissions: %w", err)
	}
	return nil
}

// Delete removes a role by ID
func (r *RoleRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Role{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}
	return nil
}

// GetAllPermissions retrieves all available permissions
func (r *RoleRepository) GetAllPermissions(ctx context.Context) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.WithContext(ctx).Order("module, name").Find(&permissions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query permissions: %w", err)
	}
	return permissions, nil
}
