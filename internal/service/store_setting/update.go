package storesetting

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// UpdateAll updates all store settings at once
func (s *Service) UpdateAll(ctx context.Context, req dto.StoreSettingsRequest) error {
	settings := req.ToMap()
	if err := s.settingRepo.UpsertAll(ctx, settings); err != nil {
		return fmt.Errorf("service: failed to update store settings: %w", err)
	}
	return nil
}
