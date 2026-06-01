package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new role
func (r *Repository) Create(ctx context.Context, rl *model.Role) error {
	if err := r.db.WithContext(ctx).Create(rl).Error; err != nil {
		return fmt.Errorf("failed to create role: %w", err)
	}
	return nil
}
