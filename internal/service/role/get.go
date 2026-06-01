package role

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// GetAllRoles retrieves all roles
func (s *Service) GetAllRoles(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get roles: %w", err)
	}

	return s.toRoleResponses(roles), nil
}

// GetAllRolesPaginated retrieves roles with pagination and optional search
func (s *Service) GetAllRolesPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.RoleResponse, dto.PaginationResponse, error) {
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

func (s *Service) toRoleResponses(roles []model.Role) []dto.RoleResponse {
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
func (s *Service) GetAllPermissions(ctx context.Context) ([]dto.PermissionResponse, error) {
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
