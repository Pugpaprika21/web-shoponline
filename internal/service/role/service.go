package role

import "web-shoponline/internal/repository"

// Service handles business logic for roles
type Service struct {
	roleRepo repository.RoleRepositoryInterface
}

// NewService creates a new role Service
func NewService(roleRepo repository.RoleRepositoryInterface) *Service {
	return &Service{roleRepo: roleRepo}
}
