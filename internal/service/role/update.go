package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// UpdateRole updates an existing role
func (s *Service) UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) error {
	rl, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}

	if req.Name != "" {
		rl.Name = req.Name
	}
	if req.Description != "" {
		rl.Description = req.Description
	}
	if req.IsActive != nil {
		rl.IsActive = *req.IsActive
	}

	if err := s.roleRepo.Update(ctx, rl); err != nil {
		return fmt.Errorf("service: failed to update role: %w", err)
	}

	// Update permissions if provided
	if len(req.PermissionIDs) > 0 {
		permissions := make([]model.Permission, len(req.PermissionIDs))
		for i, pid := range req.PermissionIDs {
			permissions[i] = model.Permission{ID: pid}
		}
		if err := s.roleRepo.UpdatePermissions(ctx, id, permissions); err != nil {
			return fmt.Errorf("service: failed to update permissions: %w", err)
		}
	}

	return nil
}
