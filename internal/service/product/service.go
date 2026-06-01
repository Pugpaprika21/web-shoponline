package product

import "web-shoponline/internal/repository"

// Service handles business logic for products
type Service struct {
	productRepo  repository.ProductRepositoryInterface
	categoryRepo repository.CategoryRepositoryInterface
}

// NewService creates a new product Service
func NewService(productRepo repository.ProductRepositoryInterface, categoryRepo repository.CategoryRepositoryInterface) *Service {
	return &Service{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}
