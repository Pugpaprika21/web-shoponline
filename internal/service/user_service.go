package service

import (
	"context"
	"fmt"
	"math"

	"golang.org/x/crypto/bcrypt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	"web-shoponline/internal/repository"
)

// UserService handles business logic for user management
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetAllPaginated retrieves users with pagination and optional search
func (s *UserService) GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.AdminUserResponse, dto.PaginationResponse, error) {
	page.Normalize()
	users, total, err := s.userRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get users: %w", err)
	}

	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return s.toUserResponses(users), pagination, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.AdminUserResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		RoleID:       &req.RoleID,
		IsActive:     true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("service: failed to create user: %w", err)
	}

	// Reload user with role
	user, err = s.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to reload user: %w", err)
	}

	resp := s.toUserResponse(*user)
	return &resp, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.RoleID != "" {
		user.RoleID = &req.RoleID
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("service: failed to update user: %w", err)
	}
	return nil
}

func (s *UserService) toUserResponses(users []model.User) []dto.AdminUserResponse {
	responses := make([]dto.AdminUserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, s.toUserResponse(u))
	}
	return responses
}

func (s *UserService) toUserResponse(u model.User) dto.AdminUserResponse {
	roleName := ""
	if u.Role != nil {
		roleName = u.Role.Name
	}
	roleID := ""
	if u.RoleID != nil {
		roleID = *u.RoleID
	}
	return dto.AdminUserResponse{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		RoleID:    roleID,
		RoleName:  roleName,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04"),
	}
}
