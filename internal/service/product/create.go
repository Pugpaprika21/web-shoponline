package product

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	productRepo "web-shoponline/internal/repository/product"
)

// CreateProduct creates a new product
func (s *Service) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	slug := productRepo.GenerateSlug(req.Name)

	p := &model.Product{
		ShopID:        req.ShopID,
		Name:          req.Name,
		Slug:          slug,
		Description:   req.Description,
		Price:         req.Price,
		StockQuantity: req.Stock,
		ImageURL:      req.ImageURL,
		IsActive:      req.IsActive,
	}

	if req.CategoryID != "" {
		p.CategoryID = &req.CategoryID
	}

	if err := s.productRepo.Create(ctx, p); err != nil {
		return nil, fmt.Errorf("service: failed to create product: %w", err)
	}

	resp := s.toProductResponse(*p)
	return &resp, nil
}
