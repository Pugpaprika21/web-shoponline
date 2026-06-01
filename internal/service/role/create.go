package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// CreateRole creates a new role
func (s *Service) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	rl := &model.Role{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if err := s.roleRepo.Create(ctx, rl); err != nil {
		return nil, fmt.Errorf("service: failed to create role: %w", err)
	}

	// Set permissions if provided
	if len(req.PermissionIDs) > 0 {
		permissions := make([]model.Permission, len(req.PermissionIDs))
		for i, id := range req.PermissionIDs {
			permissions[i] = model.Permission{ID: id}
		}
		if err := s.roleRepo.UpdatePermissions(ctx, rl.ID, permissions); err != nil {
			return nil, fmt.Errorf("service: failed to set permissions: %w", err)
		}
	}

	return &dto.RoleResponse{
		ID:          rl.ID,
		Name:        rl.Name,
		Description: rl.Description,
		IsActive:    rl.IsActive,
	}, nil
}
