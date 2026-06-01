package menu

import "web-shoponline/internal/repository"

// Service handles business logic for menus
type Service struct {
	menuRepo repository.MenuRepositoryInterface
}

// NewService creates a new menu Service
func NewService(menuRepo repository.MenuRepositoryInterface) *Service {
	return &Service{menuRepo: menuRepo}
}
