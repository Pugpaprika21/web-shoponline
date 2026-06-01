package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Update modifies an existing role
func (r *Repository) Update(ctx context.Context, rl *model.Role) error {
	if err := r.db.WithContext(ctx).Save(rl).Error; err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}
	return nil
}

// UpdatePermissions replaces the permissions for a role
func (r *Repository) UpdatePermissions(ctx context.Context, roleID string, permissions []model.Permission) error {
	var rl model.Role
	if err := r.db.WithContext(ctx).Where("id = ?", roleID).First(&rl).Error; err != nil {
		return fmt.Errorf("failed to find role: %w", err)
	}
	if err := r.db.WithContext(ctx).Model(&rl).Association("Permissions").Replace(permissions); err != nil {
		return fmt.Errorf("failed to update permissions: %w", err)
	}
	return nil
}
