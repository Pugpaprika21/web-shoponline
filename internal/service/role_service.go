package service

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	"web-shoponline/internal/repository"
)

// RoleService handles business logic for roles
type RoleService struct {
	roleRepo *repository.RoleRepository
}

// NewRoleService creates a new RoleService
func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

// GetAllRoles retrieves all roles
func (s *RoleService) GetAllRoles(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get roles: %w", err)
	}

	return s.toRoleResponses(roles), nil
}

// GetAllRolesPaginated retrieves roles with pagination and optional search
func (s *RoleService) GetAllRolesPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.RoleResponse, dto.PaginationResponse, error) {
	page.Normalize()
	roles, total, err := s.roleRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get roles: %w", err)
	}

	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return s.toRoleResponses(roles), pagination, nil
}

func (s *RoleService) toRoleResponses(roles []model.Role) []dto.RoleResponse {
	responses := make([]dto.RoleResponse, 0, len(roles))
	for _, r := range roles {
		resp := dto.RoleResponse{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			IsActive:    r.IsActive,
			Permissions: make([]dto.PermissionResponse, 0),
		}
		for _, p := range r.Permissions {
			resp.Permissions = append(resp.Permissions, dto.PermissionResponse{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Module:      p.Module,
			})
		}
		responses = append(responses, resp)
	}
	return responses
}

// GetAllPermissions retrieves all available permissions
func (s *RoleService) GetAllPermissions(ctx context.Context) ([]dto.PermissionResponse, error) {
	permissions, err := s.roleRepo.GetAllPermissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get permissions: %w", err)
	}

	responses := make([]dto.PermissionResponse, 0, len(permissions))
	for _, p := range permissions {
		responses = append(responses, dto.PermissionResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Module:      p.Module,
		})
	}

	return responses, nil
}

// CreateRole creates a new role
func (s *RoleService) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	role := &model.Role{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if err := s.roleRepo.Create(ctx, role); err != nil {
		return nil, fmt.Errorf("service: failed to create role: %w", err)
	}

	// Set permissions if provided
	if len(req.PermissionIDs) > 0 {
		permissions := make([]model.Permission, len(req.PermissionIDs))
		for i, id := range req.PermissionIDs {
			permissions[i] = model.Permission{ID: id}
		}
		if err := s.roleRepo.UpdatePermissions(ctx, role.ID, permissions); err != nil {
			return nil, fmt.Errorf("service: failed to set permissions: %w", err)
		}
	}

	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		IsActive:    role.IsActive,
	}, nil
}

// UpdateRole updates an existing role
func (s *RoleService) UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) error {
	role, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}

	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.IsActive != nil {
		role.IsActive = *req.IsActive
	}

	if err := s.roleRepo.Update(ctx, role); err != nil {
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

// DeleteRole removes a role
func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	return s.roleRepo.Delete(ctx, id)
}
