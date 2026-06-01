package category

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Delete removes a category by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Category{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
