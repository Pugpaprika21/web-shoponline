package role

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// Delete removes a role by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&model.Role{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}
	return nil
}
