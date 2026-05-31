package service

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/repository"
)

// StoreSettingService handles business logic for store settings
type StoreSettingService struct {
	settingRepo *repository.StoreSettingRepository
}

// NewStoreSettingService creates a new StoreSettingService
func NewStoreSettingService(settingRepo *repository.StoreSettingRepository) *StoreSettingService {
	return &StoreSettingService{settingRepo: settingRepo}
}

// GetAll retrieves all store settings
func (s *StoreSettingService) GetAll(ctx context.Context) (dto.StoreSettingsResponse, error) {
	settings, err := s.settingRepo.GetAll(ctx)
	if err != nil {
		return dto.StoreSettingsResponse{}, fmt.Errorf("service: failed to get store settings: %w", err)
	}
	return dto.StoreSettingsFromMap(settings), nil
}

// GetAllAsMap retrieves all store settings as a map (for template injection)
func (s *StoreSettingService) GetAllAsMap(ctx context.Context) (map[string]string, error) {
	settings, err := s.settingRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get store settings: %w", err)
	}
	return settings, nil
}

// UpdateAll updates all store settings at once
func (s *StoreSettingService) UpdateAll(ctx context.Context, req dto.StoreSettingsRequest) error {
	settings := req.ToMap()
	if err := s.settingRepo.UpsertAll(ctx, settings); err != nil {
		return fmt.Errorf("service: failed to update store settings: %w", err)
	}
	return nil
}
