package storesetting

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetAll retrieves all store settings as a map
func (r *Repository) GetAll(ctx context.Context) (map[string]string, error) {
	var settings []model.StoreSetting
	err := r.db.WithContext(ctx).Find(&settings).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query store settings: %w", err)
	}

	result := make(map[string]string)
	for _, s := range settings {
		result[s.Key] = s.Value
	}
	return result, nil
}

// GetByKey retrieves a single store setting by key
func (r *Repository) GetByKey(ctx context.Context, key string) (string, error) {
	var setting model.StoreSetting
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error
	if err != nil {
		return "", fmt.Errorf("failed to get store setting: %w", err)
	}
	return setting.Value, nil
}
