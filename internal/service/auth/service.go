package auth

import "web-shoponline/internal/repository"

// Service handles authentication business logic
type Service struct {
	userRepo repository.UserRepositoryInterface
	roleRepo repository.RoleRepositoryInterface
}

// NewService creates a new auth Service
func NewService(userRepo repository.UserRepositoryInterface, roleRepo repository.RoleRepositoryInterface) *Service {
	return &Service{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
