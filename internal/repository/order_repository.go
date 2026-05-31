package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// OrderRepository handles database operations for orders
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// GetAll retrieves all orders with user info
func (r *OrderRepository) GetAll(ctx context.Context) ([]model.Order, error) {
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
func (r *OrderRepository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Order, int64, error) {
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
func (r *OrderRepository) GetByID(ctx context.Context, id string) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Items.Product").
		Where("id = ?", id).
		First(&order).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}
	return &order, nil
}

// GetByUserID retrieves all orders for a specific user
func (r *OrderRepository) GetByUserID(ctx context.Context, userID string) ([]model.Order, error) {
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

// Create inserts a new order with items
func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	if err := r.db.WithContext(ctx).Create(order).Error; err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}

// UpdateStatus updates the status of an order
func (r *OrderRepository) UpdateStatus(ctx context.Context, id, status string) error {
	err := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}
	return nil
}
