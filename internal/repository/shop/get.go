package shop

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetByID retrieves a shop by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.WithContext(ctx).
		Preload("Owner").
		Where("id = ?", id).
		First(&shop).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get shop by id: %w", err)
	}
	return &shop, nil
}

// GetByOwnerID retrieves a shop by owner user ID
func (r *Repository) GetByOwnerID(ctx context.Context, ownerID string) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.WithContext(ctx).
		Preload("Owner").
		Where("owner_id = ?", ownerID).
		First(&shop).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get shop by owner id: %w", err)
	}
	return &shop, nil
}

// GetBySlug retrieves a shop by slug
func (r *Repository) GetBySlug(ctx context.Context, slug string) (*model.Shop, error) {
	var shop model.Shop
	err := r.db.WithContext(ctx).
		Preload("Owner").
		Where("slug = ? AND is_active = ?", slug, true).
		First(&shop).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get shop by slug: %w", err)
	}
	return &shop, nil
}

// GetAll retrieves all active shops
func (r *Repository) GetAll(ctx context.Context) ([]model.Shop, error) {
	var shops []model.Shop
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Order("name").
		Find(&shops).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query shops: %w", err)
	}
	return shops, nil
}
