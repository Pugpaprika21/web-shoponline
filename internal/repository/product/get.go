package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all active products with category and shop
func (r *Repository) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Where("is_active = ?", true).
		Order("created_at DESC").
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	return products, nil
}

// GetAllPaginated retrieves active products with pagination and optional search
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Product{}).Where("is_active = ?", true).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Where("is_active = ?", true).
		Where("deleted_at IS NULL")
	if search != "" {
		q = q.Where("name ILIKE ?", "%"+search+"%")
	}
	err := q.Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&products).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query products: %w", err)
	}
	return products, total, nil
}

// GetAllAdmin retrieves all products (including inactive) for admin
func (r *Repository) GetAllAdmin(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Order("created_at DESC").
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	return products, nil
}

// GetAllAdminPaginated retrieves all products with pagination and optional search for admin
func (r *Repository) GetAllAdminPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Product{}).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Where("deleted_at IS NULL")
	if search != "" {
		q = q.Where("name ILIKE ?", "%"+search+"%")
	}
	err := q.Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&products).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query products: %w", err)
	}
	return products, total, nil
}

// GetByID retrieves a product by its ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}
	return &product, nil
}

// GetByShopID retrieves all active products for a specific shop
func (r *Repository) GetByShopID(ctx context.Context, shopID string) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Shop").
		Where("shop_id = ? AND is_active = ?", shopID, true).
		Order("created_at DESC").
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get products by shop: %w", err)
	}
	return products, nil
}
