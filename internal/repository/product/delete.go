package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Delete soft-deletes a product by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Product{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
