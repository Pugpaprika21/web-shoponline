package cart

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// GetOrCreateByUserID retrieves or creates a cart for a user
func (r *Repository) GetOrCreateByUserID(ctx context.Context, userID string) (*model.Cart, error) {
	var c model.Cart
	err := r.db.WithContext(ctx).
		Preload("Items.Product").
		Where("user_id = ?", userID).
		First(&c).Error

	if err == gorm.ErrRecordNotFound {
		c = model.Cart{UserID: userID}
		if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
			return nil, fmt.Errorf("failed to create cart: %w", err)
		}
		return &c, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	return &c, nil
}

// GetItemCount returns the total number of items in a cart
func (r *Repository) GetItemCount(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.CartItem{}).
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count cart items: %w", err)
	}
	return count, nil
}
