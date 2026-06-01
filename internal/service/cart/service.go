package cart

import "web-shoponline/internal/repository"

// Service handles business logic for shopping carts
type Service struct {
	cartRepo    repository.CartRepositoryInterface
	productRepo repository.ProductRepositoryInterface
}

// NewService creates a new cart Service
func NewService(cartRepo repository.CartRepositoryInterface, productRepo repository.ProductRepositoryInterface) *Service {
	return &Service{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}
