package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Update modifies an existing menu
func (r *Repository) Update(ctx context.Context, m *model.Menu) error {
	if err := r.db.WithContext(ctx).Save(m).Error; err != nil {
		return fmt.Errorf("failed to update menu: %w", err)
	}
	return nil
}
