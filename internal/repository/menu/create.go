package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new menu
func (r *Repository) Create(ctx context.Context, m *model.Menu) error {
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return fmt.Errorf("failed to create menu: %w", err)
	}
	return nil
}
