package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// CreateMenu creates a new menu
func (s *Service) CreateMenu(ctx context.Context, req dto.CreateMenuRequest) (*dto.MenuResponse, error) {
	m := &model.Menu{
		Name:      req.Name,
		Icon:      req.Icon,
		URL:       req.URL,
		SortOrder: req.SortOrder,
		IsActive:  req.IsActive,
	}
	if req.ParentID != "" {
		m.ParentID = &req.ParentID
	}

	if err := s.menuRepo.Create(ctx, m); err != nil {
		return nil, fmt.Errorf("service: failed to create menu: %w", err)
	}

	return &dto.MenuResponse{
		ID:        m.ID,
		Name:      m.Name,
		Icon:      m.Icon,
		URL:       m.URL,
		SortOrder: m.SortOrder,
		IsActive:  m.IsActive,
	}, nil
}
