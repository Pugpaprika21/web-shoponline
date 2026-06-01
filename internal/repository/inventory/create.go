package inventory

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// Create inserts a new inventory record and updates product stock
func (r *Repository) Create(ctx context.Context, record *model.Inventory) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create inventory record
		if err := tx.Create(record).Error; err != nil {
			return fmt.Errorf("failed to create inventory record: %w", err)
		}

		// Update product stock quantity
		var adjustment int
		switch record.Type {
		case "in":
			adjustment = record.Quantity
		case "out":
			adjustment = -record.Quantity
		case "adjustment":
			adjustment = record.Quantity // can be negative
		}

		err := tx.Model(&model.Product{}).
			Where("id = ?", record.ProductID).
			Update("stock_quantity", gorm.Expr("stock_quantity + ?", adjustment)).Error
		if err != nil {
			return fmt.Errorf("failed to update product stock: %w", err)
		}

		return nil
	})
}
