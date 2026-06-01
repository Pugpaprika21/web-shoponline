package shop

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Create inserts a new shop
func (r *Repository) Create(ctx context.Context, shop *model.Shop) error {
	if err := r.db.WithContext(ctx).Create(shop).Error; err != nil {
		return fmt.Errorf("failed to create shop: %w", err)
	}
	return nil
}

// Update modifies an existing shop
func (r *Repository) Update(ctx context.Context, shop *model.Shop) error {
	if err := r.db.WithContext(ctx).Save(shop).Error; err != nil {
		return fmt.Errorf("failed to update shop: %w", err)
	}
	return nil
}
