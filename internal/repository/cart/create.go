package cart

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// AddItem adds or updates an item in the cart
func (r *Repository) AddItem(ctx context.Context, cartID, productID string, quantity int) error {
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
