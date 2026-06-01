package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	productRepo "web-shoponline/internal/repository/product"
)

// UpdateProduct updates an existing product
func (s *Service) UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	p, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: product not found: %w", err)
	}

	if req.Name != "" {
		p.Name = req.Name
		p.Slug = productRepo.GenerateSlug(req.Name)
	}
	if req.CategoryID != "" {
		p.CategoryID = &req.CategoryID
	}
	if req.Description != "" {
		p.Description = req.Description
	}
	if req.Price > 0 {
		p.Price = req.Price
	}
	if req.Stock > 0 {
		p.StockQuantity = req.Stock
	}
	if req.ImageURL != "" {
		p.ImageURL = req.ImageURL
	}
	if req.IsActive != nil {
		p.IsActive = *req.IsActive
	}

	if err := s.productRepo.Update(ctx, p); err != nil {
		return nil, fmt.Errorf("service: failed to update product: %w", err)
	}

	resp := s.toProductResponse(*p)
	return &resp, nil
}
