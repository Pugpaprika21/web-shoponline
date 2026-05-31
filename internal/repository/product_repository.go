package repository

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll retrieves all active products with category
func (r *ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Where("is_active = ?", true).
		Order("created_at DESC").
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	return products, nil
}

// GetAllPaginated retrieves active products with pagination and optional search
func (r *ProductRepository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Product{}).Where("is_active = ?", true).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Category").
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
func (r *ProductRepository) GetAllAdmin(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Order("created_at DESC").
		Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	return products, nil
}

// GetAllAdminPaginated retrieves all products with pagination and optional search for admin
func (r *ProductRepository) GetAllAdminPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Product{}).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Category").
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
func (r *ProductRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}
	return &product, nil
}

// Create inserts a new product
func (r *ProductRepository) Create(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

// Update modifies an existing product
func (r *ProductRepository) Update(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Save(product).Error; err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

// Delete soft-deletes a product by ID
func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Product{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

// GenerateSlug creates a URL-friendly slug from a product name
func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
