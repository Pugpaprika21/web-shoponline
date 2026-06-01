package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// CreateUser creates a new user
func (s *Service) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.AdminUserResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Get roles by IDs
	roles, err := s.getRolesByIDs(ctx, req.RoleIDs)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		Roles:        roles,
		IsActive:     true,
	}

	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("service: failed to create user: %w", err)
	}

	// Reload user with roles
	u, err = s.userRepo.GetByID(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to reload user: %w", err)
	}

	resp := s.toUserResponse(*u)
	return &resp, nil
}

func (s *Service) getRolesByIDs(ctx context.Context, roleIDs []string) ([]model.Role, error) {
	allRoles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	roleMap := make(map[string]model.Role)
	for _, r := range allRoles {
		roleMap[r.ID] = r
	}

	var roles []model.Role
	for _, id := range roleIDs {
		if rl, ok := roleMap[id]; ok {
			roles = append(roles, rl)
		}
	}
	return roles, nil
}
