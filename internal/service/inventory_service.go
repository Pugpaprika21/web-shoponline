package service

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	"web-shoponline/internal/repository"
)

// InventoryService handles business logic for inventory
type InventoryService struct {
	inventoryRepo *repository.InventoryRepository
	productRepo   *repository.ProductRepository
}

// NewInventoryService creates a new InventoryService
func NewInventoryService(inventoryRepo *repository.InventoryRepository, productRepo *repository.ProductRepository) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
		productRepo:   productRepo,
	}
}

// GetAll retrieves all inventory records
func (s *InventoryService) GetAll(ctx context.Context) ([]dto.InventoryResponse, error) {
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
func (s *InventoryService) GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.InventoryResponse, dto.PaginationResponse, error) {
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

// AddInventory creates a new inventory record
func (s *InventoryService) AddInventory(ctx context.Context, req dto.CreateInventoryRequest, createdBy string) (*dto.InventoryResponse, error) {
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
