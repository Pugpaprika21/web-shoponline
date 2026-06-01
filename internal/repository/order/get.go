package order

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all orders with user info
func (r *Repository) GetAll(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Order("created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}
	return orders, nil
}

// GetAllPaginated retrieves orders with pagination and optional search
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Order{}).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("id::text ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("User").
		Where("deleted_at IS NULL")
	if search != "" {
		q = q.Where("id::text ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	err := q.Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&orders).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query orders: %w", err)
	}
	return orders, total, nil
}

// GetByID retrieves an order by its ID with items
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Order, error) {
	var o model.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Items.Product").
		Where("id = ?", id).
		First(&o).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}
	return &o, nil
}

// GetByUserID retrieves all orders for a specific user
func (r *Repository) GetByUserID(ctx context.Context, userID string) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by user: %w", err)
	}
	return orders, nil
}

// GetByShopID retrieves all orders that contain products from a specific shop
func (r *Repository) GetByShopID(ctx context.Context, shopID string) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Items.Product").
		Joins("JOIN order_items ON order_items.order_id = orders.id").
		Joins("JOIN products ON products.id = order_items.product_id").
		Where("products.shop_id = ?", shopID).
		Group("orders.id").
		Order("orders.created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by shop: %w", err)
	}
	return orders, nil
}
