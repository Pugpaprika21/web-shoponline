package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// InventoryRepository handles database operations for inventory
type InventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository creates a new InventoryRepository
func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

// GetAll retrieves all inventory records
func (r *InventoryRepository) GetAll(ctx context.Context) ([]model.Inventory, error) {
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
func (r *InventoryRepository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Inventory, int64, error) {
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
func (r *InventoryRepository) GetByProductID(ctx context.Context, productID string) ([]model.Inventory, error) {
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

// Create inserts a new inventory record and updates product stock
func (r *InventoryRepository) Create(ctx context.Context, record *model.Inventory) error {
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
