package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// CartRepository handles database operations for shopping carts
type CartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new CartRepository
func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

// GetOrCreateByUserID retrieves or creates a cart for a user
func (r *CartRepository) GetOrCreateByUserID(ctx context.Context, userID string) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.WithContext(ctx).
		Preload("Items.Product").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err == gorm.ErrRecordNotFound {
		cart = model.Cart{UserID: userID}
		if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
			return nil, fmt.Errorf("failed to create cart: %w", err)
		}
		return &cart, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	return &cart, nil
}

// AddItem adds or updates an item in the cart
func (r *CartRepository) AddItem(ctx context.Context, cartID, productID string, quantity int) error {
	var item model.CartItem
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&item).Error

	if err == gorm.ErrRecordNotFound {
		// Create new item
		item = model.CartItem{
			CartID:    cartID,
			ProductID: productID,
			Quantity:  quantity,
		}
		if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
			return fmt.Errorf("failed to add cart item: %w", err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to check cart item: %w", err)
	}

	// Update quantity
	item.Quantity += quantity
	if err := r.db.WithContext(ctx).Save(&item).Error; err != nil {
		return fmt.Errorf("failed to update cart item: %w", err)
	}
	return nil
}

// RemoveItem removes an item from the cart
func (r *CartRepository) RemoveItem(ctx context.Context, cartID, itemID string) error {
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND id = ?", cartID, itemID).
		Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("failed to remove cart item: %w", err)
	}
	return nil
}

// ClearCart removes all items from a cart
func (r *CartRepository) ClearCart(ctx context.Context, cartID string) error {
	err := r.db.WithContext(ctx).
		Where("cart_id = ?", cartID).
		Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}
	return nil
}

// GetItemCount returns the total number of items in a cart
func (r *CartRepository) GetItemCount(ctx context.Context, userID string) (int64, error) {
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
