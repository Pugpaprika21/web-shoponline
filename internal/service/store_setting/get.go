package storesetting

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// GetAll retrieves all store settings
func (s *Service) GetAll(ctx context.Context) (dto.StoreSettingsResponse, error) {
	settings, err := s.settingRepo.GetAll(ctx)
	if err != nil {
		return dto.StoreSettingsResponse{}, fmt.Errorf("service: failed to get store settings: %w", err)
	}
	return dto.StoreSettingsFromMap(settings), nil
}

// GetAllAsMap retrieves all store settings as a map (for template injection)
func (s *Service) GetAllAsMap(ctx context.Context) (map[string]string, error) {
	settings, err := s.settingRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get store settings: %w", err)
	}
	return settings, nil
}
