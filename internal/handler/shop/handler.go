package shop

import (
	"web-shoponline/internal/repository"
)

// Handler handles HTTP requests for shop pages
type Handler struct {
	shopRepo    repository.ShopRepositoryInterface
	productRepo repository.ProductRepositoryInterface
}

// NewHandler creates a new shop Handler
func NewHandler(shopRepo repository.ShopRepositoryInterface, productRepo repository.ProductRepositoryInterface) *Handler {
	return &Handler{shopRepo: shopRepo, productRepo: productRepo}
}
