package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"web-shoponline/internal/dto"
)

// Login authenticates a user and returns user data
func (s *Service) Login(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	roleNames := getRoleNames(user)

	resp := &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		RoleNames: roleNames,
		IsActive:  user.IsActive,
	}

	return resp, nil
}
