package service

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	"web-shoponline/internal/repository"
)

// MenuService handles business logic for menus
type MenuService struct {
	menuRepo *repository.MenuRepository
}

// NewMenuService creates a new MenuService
func NewMenuService(menuRepo *repository.MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}

// GetAllMenus retrieves all menus as a tree structure
func (s *MenuService) GetAllMenus(ctx context.Context) ([]dto.MenuResponse, error) {
	menus, err := s.menuRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get menus: %w", err)
	}

	return s.buildMenuTree(menus), nil
}

// GetAllMenusPaginated retrieves menus with pagination and optional search (flat list)
func (s *MenuService) GetAllMenusPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.MenuResponse, dto.PaginationResponse, error) {
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

func (s *MenuService) buildMenuTree(menus []model.Menu) []dto.MenuResponse {
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

// CreateMenu creates a new menu
func (s *MenuService) CreateMenu(ctx context.Context, req dto.CreateMenuRequest) (*dto.MenuResponse, error) {
	menu := &model.Menu{
		Name:      req.Name,
		Icon:      req.Icon,
		URL:       req.URL,
		SortOrder: req.SortOrder,
		IsActive:  req.IsActive,
	}
	if req.ParentID != "" {
		menu.ParentID = &req.ParentID
	}

	if err := s.menuRepo.Create(ctx, menu); err != nil {
		return nil, fmt.Errorf("service: failed to create menu: %w", err)
	}

	return &dto.MenuResponse{
		ID:        menu.ID,
		Name:      menu.Name,
		Icon:      menu.Icon,
		URL:       menu.URL,
		SortOrder: menu.SortOrder,
		IsActive:  menu.IsActive,
	}, nil
}

// UpdateMenu updates an existing menu
func (s *MenuService) UpdateMenu(ctx context.Context, id string, req dto.UpdateMenuRequest) error {
	menu, err := s.menuRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("menu not found: %w", err)
	}

	if req.Name != "" {
		menu.Name = req.Name
	}
	if req.Icon != "" {
		menu.Icon = req.Icon
	}
	if req.URL != "" {
		menu.URL = req.URL
	}
	if req.SortOrder > 0 {
		menu.SortOrder = req.SortOrder
	}
	if req.IsActive != nil {
		menu.IsActive = *req.IsActive
	}
	if req.ParentID != "" {
		menu.ParentID = &req.ParentID
	}

	return s.menuRepo.Update(ctx, menu)
}

// DeleteMenu removes a menu
func (s *MenuService) DeleteMenu(ctx context.Context, id string) error {
	return s.menuRepo.Delete(ctx, id)
}
