package user

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// GetAllPaginated retrieves users with pagination and optional search
func (s *Service) GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.AdminUserResponse, dto.PaginationResponse, error) {
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

func (s *Service) toUserResponses(users []model.User) []dto.AdminUserResponse {
	responses := make([]dto.AdminUserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, s.toUserResponse(u))
	}
	return responses
}

func (s *Service) toUserResponse(u model.User) dto.AdminUserResponse {
	roleIDs := make([]string, 0, len(u.Roles))
	roleNames := make([]string, 0, len(u.Roles))
	for _, r := range u.Roles {
		roleIDs = append(roleIDs, r.ID)
		roleNames = append(roleNames, r.Name)
	}
	return dto.AdminUserResponse{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		RoleIDs:   roleIDs,
		RoleNames: roleNames,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04"),
	}
}
