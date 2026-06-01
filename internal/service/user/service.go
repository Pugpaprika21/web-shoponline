package user

import "web-shoponline/internal/repository"

// Service handles business logic for user management
type Service struct {
	userRepo repository.UserRepositoryInterface
	roleRepo repository.RoleRepositoryInterface
}

// NewService creates a new user Service
func NewService(userRepo repository.UserRepositoryInterface, roleRepo repository.RoleRepositoryInterface) *Service {
	return &Service{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
