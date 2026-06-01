package inventory

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all inventory records
func (r *Repository) GetAll(ctx context.Context) ([]model.Inventory, error) {
	var records []model.Inventory
	err := r.db.WithContext(ctx).
		Preload("Product").
		Order("created_at DESC").
		Find(&records).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query inventory: %w", err)
	}
	return records, nil
}

// GetAllPaginated retrieves inventory records with pagination and optional search
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Inventory, int64, error) {
	var records []model.Inventory
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Inventory{})
	if search != "" {
		query = query.Joins("JOIN products ON products.id = inventories.product_id").
			Where("products.name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Product")
	if search != "" {
		q = q.Joins("JOIN products ON products.id = inventories.product_id").
			Where("products.name ILIKE ?", "%"+search+"%")
	}
	err := q.Order("inventories.created_at DESC").
		Offset(offset).Limit(limit).
		Find(&records).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query inventory: %w", err)
	}
	return records, total, nil
}

// GetByProductID retrieves inventory records for a specific product
func (r *Repository) GetByProductID(ctx context.Context, productID string) ([]model.Inventory, error) {
	var records []model.Inventory
	err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Order("created_at DESC").
		Find(&records).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query inventory by product: %w", err)
	}
	return records, nil
}
