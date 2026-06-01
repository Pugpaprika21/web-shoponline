package inventory

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// AddInventory creates a new inventory record
func (s *Service) AddInventory(ctx context.Context, req dto.CreateInventoryRequest, createdBy string) (*dto.InventoryResponse, error) {
	// Verify product exists
	product, err := s.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	record := &model.Inventory{
		ProductID: req.ProductID,
		Type:      req.Type,
		Quantity:  req.Quantity,
		Note:      req.Note,
		CreatedBy: createdBy,
	}

	if err := s.inventoryRepo.Create(ctx, record); err != nil {
		return nil, fmt.Errorf("service: failed to add inventory: %w", err)
	}

	return &dto.InventoryResponse{
		ID:          record.ID,
		ProductID:   record.ProductID,
		ProductName: product.Name,
		Type:        record.Type,
		Quantity:    record.Quantity,
		Note:        record.Note,
		CreatedAt:   record.CreatedAt,
	}, nil
}
