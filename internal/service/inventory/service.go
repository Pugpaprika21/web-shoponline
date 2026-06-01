package inventory

import "web-shoponline/internal/repository"

// Service handles business logic for inventory
type Service struct {
	inventoryRepo repository.InventoryRepositoryInterface
	productRepo   repository.ProductRepositoryInterface
}

// NewService creates a new inventory Service
func NewService(inventoryRepo repository.InventoryRepositoryInterface, productRepo repository.ProductRepositoryInterface) *Service {
	return &Service{
		inventoryRepo: inventoryRepo,
		productRepo:   productRepo,
	}
}
