package user

import (
	"context"
	"fmt"

	"web-shoponline/internal/model"
)

// GetByEmail retrieves a user by email
func (r *Repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var u model.User
	err := r.db.WithContext(ctx).
		Preload("Roles.Permissions").
		Where("email = ?", email).
		First(&u).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &u, nil
}

// GetByID retrieves a user by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var u model.User
	err := r.db.WithContext(ctx).
		Preload("Roles.Permissions").
		Where("id = ?", id).
		First(&u).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &u, nil
}

// GetAll retrieves all users
func (r *Repository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).
		Preload("Roles").
		Order("created_at DESC").
		Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	return users, nil
}

// GetAllPaginated retrieves users with pagination and optional search
func (r *Repository) GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.WithContext(ctx).Model(&model.User{})
	if search != "" {
		query = query.Where("email ILIKE ? OR full_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&total)

	q := r.db.WithContext(ctx).
		Preload("Roles")
	if search != "" {
		q = q.Where("email ILIKE ? OR full_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	err := q.Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query users: %w", err)
	}
	return users, total, nil
}
