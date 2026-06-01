package order

import "web-shoponline/internal/repository"

// Service handles business logic for orders
type Service struct {
	orderRepo repository.OrderRepositoryInterface
}

// NewService creates a new order Service
func NewService(orderRepo repository.OrderRepositoryInterface) *Service {
	return &Service{orderRepo: orderRepo}
}
