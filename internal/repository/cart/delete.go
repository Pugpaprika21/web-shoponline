package cart

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// RemoveItem removes an item from the cart
func (r *Repository) RemoveItem(ctx context.Context, cartID, itemID string) error {
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND id = ?", cartID, itemID).
		Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("failed to remove cart item: %w", err)
	}
	return nil
}

// ClearCart removes all items from a cart
func (r *Repository) ClearCart(ctx context.Context, cartID string) error {
	err := r.db.WithContext(ctx).
		Where("cart_id = ?", cartID).
		Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}
	return nil
}
