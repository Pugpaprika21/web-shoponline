package order

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new order with items
func (r *Repository) Create(ctx context.Context, o *model.Order) error {
	if err := r.db.WithContext(ctx).Create(o).Error; err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}
