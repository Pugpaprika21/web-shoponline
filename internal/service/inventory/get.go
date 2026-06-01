package inventory

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
)

// GetAll retrieves all inventory records
func (s *Service) GetAll(ctx context.Context) ([]dto.InventoryResponse, error) {
	records, err := s.inventoryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get inventory: %w", err)
	}

	responses := make([]dto.InventoryResponse, 0, len(records))
	for _, r := range records {
		responses = append(responses, dto.InventoryResponse{
			ID:          r.ID,
			ProductID:   r.ProductID,
			ProductName: r.Product.Name,
			Type:        r.Type,
			Quantity:    r.Quantity,
			Note:        r.Note,
			CreatedAt:   r.CreatedAt,
		})
	}

	return responses, nil
}

// GetAllPaginated retrieves inventory records with pagination and optional search
func (s *Service) GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.InventoryResponse, dto.PaginationResponse, error) {
	page.Normalize()
	records, total, err := s.inventoryRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get inventory: %w", err)
	}

	responses := make([]dto.InventoryResponse, 0, len(records))
	for _, r := range records {
		responses = append(responses, dto.InventoryResponse{
			ID:          r.ID,
			ProductID:   r.ProductID,
			ProductName: r.Product.Name,
			Type:        r.Type,
			Quantity:    r.Quantity,
			Note:        r.Note,
			CreatedAt:   r.CreatedAt,
		})
	}

	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return responses, pagination, nil
}
