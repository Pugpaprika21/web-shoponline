package order

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// UpdateStatus updates the status of an order
func (r *Repository) UpdateStatus(ctx context.Context, id, status string) error {
	err := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}
	return nil
}
