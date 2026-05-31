package service

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/repository"
)

// CartService handles business logic for shopping carts
type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

// NewCartService creates a new CartService
func NewCartService(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

// GetCart retrieves the user's cart
func (s *CartService) GetCart(ctx context.Context, userID string) (*dto.CartResponse, error) {
	cart, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get cart: %w", err)
	}

	resp := &dto.CartResponse{
		ID:    cart.ID,
		Items: make([]dto.CartItemResponse, 0),
	}

	var totalPrice float64
	for _, item := range cart.Items {
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

	resp.TotalItems = len(cart.Items)
	resp.TotalPrice = totalPrice

	return resp, nil
}

// AddToCart adds a product to the user's cart
func (s *CartService) AddToCart(ctx context.Context, userID string, req dto.AddToCartRequest) error {
	// Verify product exists and has stock
	product, err := s.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.StockQuantity < req.Quantity {
		return fmt.Errorf("insufficient stock (available: %d)", product.StockQuantity)
	}

	cart, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: failed to get cart: %w", err)
	}

	if err := s.cartRepo.AddItem(ctx, cart.ID, req.ProductID, req.Quantity); err != nil {
		return fmt.Errorf("service: failed to add item to cart: %w", err)
	}

	return nil
}

// RemoveFromCart removes an item from the user's cart
func (s *CartService) RemoveFromCart(ctx context.Context, userID, itemID string) error {
	cart, err := s.cartRepo.GetOrCreateByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: failed to get cart: %w", err)
	}

	if err := s.cartRepo.RemoveItem(ctx, cart.ID, itemID); err != nil {
		return fmt.Errorf("service: failed to remove item: %w", err)
	}

	return nil
}

// GetCartItemCount returns the number of items in the user's cart
func (s *CartService) GetCartItemCount(ctx context.Context, userID string) (int64, error) {
	return s.cartRepo.GetItemCount(ctx, userID)
}
