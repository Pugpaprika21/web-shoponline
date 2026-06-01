package auth

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// GetUserByID retrieves user info by ID
func (s *Service) GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
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

// HasRole checks if a user has a specific role
func HasRole(user *model.User, roleName string) bool {
	for _, r := range user.Roles {
		if r.Name == roleName {
			return true
		}
	}
	return false
}

// getRoleNames extracts role names from a user
func getRoleNames(user *model.User) []string {
	names := make([]string, 0, len(user.Roles))
	for _, r := range user.Roles {
		names = append(names, r.Name)
	}
	return names
}
