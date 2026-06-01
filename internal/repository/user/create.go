package user

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new user
func (r *Repository) Create(ctx context.Context, u *model.User) error {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
