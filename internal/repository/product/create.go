package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new product
func (r *Repository) Create(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}
