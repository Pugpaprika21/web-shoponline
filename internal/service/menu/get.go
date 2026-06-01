package menu

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// GetAllMenus retrieves all menus as a tree structure
func (s *Service) GetAllMenus(ctx context.Context) ([]dto.MenuResponse, error) {
	menus, err := s.menuRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get menus: %w", err)
	}

	return s.buildMenuTree(menus), nil
}

// GetAllMenusPaginated retrieves menus with pagination and optional search (flat list)
func (s *Service) GetAllMenusPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.MenuResponse, dto.PaginationResponse, error) {
	page.Normalize()
	menus, total, err := s.menuRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get menus: %w", err)
	}

	result := s.buildMenuTree(menus)

	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return result, pagination, nil
}

func (s *Service) buildMenuTree(menus []model.Menu) []dto.MenuResponse {
	// Build tree structure
	menuMap := make(map[string]*dto.MenuResponse)

	for _, m := range menus {
		resp := dto.MenuResponse{
			ID:        m.ID,
			Name:      m.Name,
			Icon:      m.Icon,
			URL:       m.URL,
			SortOrder: m.SortOrder,
			IsActive:  m.IsActive,
			Children:  make([]dto.MenuResponse, 0),
		}
		if m.ParentID != nil {
			resp.ParentID = *m.ParentID
		}
		menuMap[m.ID] = &resp
	}

	for _, m := range menus {
		if m.ParentID != nil {
			if parent, ok := menuMap[*m.ParentID]; ok {
				parent.Children = append(parent.Children, *menuMap[m.ID])
			}
		}
	}

	// Re-build roots with children
	result := make([]dto.MenuResponse, 0)
	for _, m := range menus {
		if m.ParentID == nil {
			item := menuMap[m.ID]
			result = append(result, *item)
		}
	}

	return result
}
