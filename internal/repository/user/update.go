package user

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Update modifies an existing user
func (r *Repository) Update(ctx context.Context, u *model.User) error {
	if err := r.db.WithContext(ctx).Save(u).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// ReplaceRoles replaces all roles for a user using the many-to-many association
func (r *Repository) ReplaceRoles(ctx context.Context, u *model.User, roles []model.Role) error {
	if err := r.db.WithContext(ctx).Model(u).Association("Roles").Replace(roles); err != nil {
		return fmt.Errorf("failed to replace user roles: %w", err)
	}
	return nil
}
