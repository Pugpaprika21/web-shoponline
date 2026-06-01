package category

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new category
func (r *Repository) Create(ctx context.Context, cat *model.Category) error {
	if err := r.db.WithContext(ctx).Create(cat).Error; err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}
