package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"web-shoponline/internal/model"
)

// StoreSettingRepository handles database operations for store settings
type StoreSettingRepository struct {
	db *gorm.DB
}

// NewStoreSettingRepository creates a new StoreSettingRepository
func NewStoreSettingRepository(db *gorm.DB) *StoreSettingRepository {
	return &StoreSettingRepository{db: db}
}

// GetAll retrieves all store settings as a map
func (r *StoreSettingRepository) GetAll(ctx context.Context) (map[string]string, error) {
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
func (r *StoreSettingRepository) GetByKey(ctx context.Context, key string) (string, error) {
	var setting model.StoreSetting
	err := r.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error
	if err != nil {
		return "", fmt.Errorf("failed to get store setting: %w", err)
	}
	return setting.Value, nil
}

// UpsertAll upserts multiple store settings
func (r *StoreSettingRepository) UpsertAll(ctx context.Context, settings map[string]string) error {
	for key, value := range settings {
		setting := model.StoreSetting{
			Key:   key,
			Value: value,
		}
		err := r.db.WithContext(ctx).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "key"}},
				DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
			}).
			Create(&setting).Error
		if err != nil {
			return fmt.Errorf("failed to upsert store setting %s: %w", key, err)
		}
	}
	return nil
}
