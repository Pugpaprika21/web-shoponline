package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Delete removes a menu by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Menu{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete menu: %w", err)
	}
	return nil
}
