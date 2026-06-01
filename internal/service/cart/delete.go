package cart

import (
	"context"
	"fmt"
)

// RemoveFromCart removes an item from the user's cart
func (s *Service) RemoveFromCart(ctx context.Context, userID, itemID string) error {
	c, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: failed to get cart: %w", err)
	}

	if err := s.cartRepo.RemoveItem(ctx, c.ID, itemID); err != nil {
		return fmt.Errorf("service: failed to remove item: %w", err)
	}

	return nil
}
