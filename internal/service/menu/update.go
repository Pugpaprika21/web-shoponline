package menu

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// UpdateMenu updates an existing menu
func (s *Service) UpdateMenu(ctx context.Context, id string, req dto.UpdateMenuRequest) error {
	m, err := s.menuRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("menu not found: %w", err)
	}

	if req.Name != "" {
		m.Name = req.Name
	}
	if req.Icon != "" {
		m.Icon = req.Icon
	}
	if req.URL != "" {
		m.URL = req.URL
	}
	if req.SortOrder > 0 {
		m.SortOrder = req.SortOrder
	}
	if req.IsActive != nil {
		m.IsActive = *req.IsActive
	}
	if req.ParentID != "" {
		m.ParentID = &req.ParentID
	}

	return s.menuRepo.Update(ctx, m)
}
