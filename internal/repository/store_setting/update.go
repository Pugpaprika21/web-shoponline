package storesetting

import (
	"context"
	"fmt"

	"gorm.io/gorm/clause"

	"web-shoponline/internal/model"
)

// UpsertAll upserts multiple store settings
func (r *Repository) UpsertAll(ctx context.Context, settings map[string]string) error {
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
