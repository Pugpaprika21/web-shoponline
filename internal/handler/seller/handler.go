package seller

import (
	"web-shoponline/internal/repository"
)

// Handler handles HTTP requests for the seller dashboard
type Handler struct {
	shopRepo    repository.ShopRepositoryInterface
	productRepo repository.ProductRepositoryInterface
	orderRepo   repository.OrderRepositoryInterface
}

// NewHandler creates a new seller Handler
func NewHandler(shopRepo repository.ShopRepositoryInterface, productRepo repository.ProductRepositoryInterface, orderRepo repository.OrderRepositoryInterface) *Handler {
	return &Handler{
		shopRepo:    shopRepo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}
