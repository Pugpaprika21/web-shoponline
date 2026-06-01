package user

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// UpdateUser updates an existing user
func (s *Service) UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) error {
	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if req.FullName != "" {
		u.FullName = req.FullName
	}
	if req.IsActive != nil {
		u.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(ctx, u); err != nil {
		return fmt.Errorf("service: failed to update user: %w", err)
	}

	// Update roles if provided
	if len(req.RoleIDs) > 0 {
		roles, err := s.getRolesByIDs(ctx, req.RoleIDs)
		if err != nil {
			return err
		}
		if err := s.userRepo.ReplaceRoles(ctx, u, roles); err != nil {
			return fmt.Errorf("service: failed to update user roles: %w", err)
		}
	}

	return nil
}
