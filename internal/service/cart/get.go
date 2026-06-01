package cart

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// GetCart retrieves the user's cart
func (s *Service) GetCart(ctx context.Context, userID string) (*dto.CartResponse, error) {
	c, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get cart: %w", err)
	}

	resp := &dto.CartResponse{
		ID:    c.ID,
		Items: make([]dto.CartItemResponse, 0),
	}

	var totalPrice float64
	for _, item := range c.Items {
		subtotal := float64(item.Quantity) * item.Product.Price
		totalPrice += subtotal

		resp.Items = append(resp.Items, dto.CartItemResponse{
			ID:          item.ID,
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			ImageURL:    item.Product.ImageURL,
			UnitPrice:   item.Product.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	resp.TotalItems = len(c.Items)
	resp.TotalPrice = totalPrice

	return resp, nil
}

// GetCartItemCount returns the number of items in the user's cart
func (s *Service) GetCartItemCount(ctx context.Context, userID string) (int64, error) {
	return s.cartRepo.GetItemCount(ctx, userID)
}
