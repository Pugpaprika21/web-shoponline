package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Update modifies an existing product
func (r *Repository) Update(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Save(product).Error; err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}
