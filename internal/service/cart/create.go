package cart

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// AddToCart adds a product to the user's cart
func (s *Service) AddToCart(ctx context.Context, userID string, req dto.AddToCartRequest) error {
	// Verify product exists and has stock
	product, err := s.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.StockQuantity < req.Quantity {
		return fmt.Errorf("insufficient stock (available: %d)", product.StockQuantity)
	}

	c, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: failed to get cart: %w", err)
	}

	if err := s.cartRepo.AddItem(ctx, c.ID, req.ProductID, req.Quantity); err != nil {
		return fmt.Errorf("service: failed to add item to cart: %w", err)
	}

	return nil
}
