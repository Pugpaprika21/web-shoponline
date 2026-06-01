package auth

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// Register creates a new customer user
func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Get customer role
	roles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	var customerRole *model.Role
	for i := range roles {
		if roles[i].Name == "customer" {
			customerRole = &roles[i]
			break
		}
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		IsActive:     true,
	}
	if customerRole != nil {
		user.Roles = []model.Role{*customerRole}
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		RoleNames: []string{"customer"},
		IsActive:  true,
	}, nil
}
