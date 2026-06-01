package product

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
)

// GetAllProducts retrieves all active products as DTOs
func (s *Service) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get products: %w", err)
	}
	return s.toProductResponses(products), nil
}

// GetAllProductsPaginated retrieves active products with pagination and optional search
func (s *Service) GetAllProductsPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error) {
	page.Normalize()
	products, total, err := s.productRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get products: %w", err)
	}
	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return s.toProductResponses(products), pagination, nil
}

// GetAllProductsAdmin retrieves all products (including inactive) as DTOs
func (s *Service) GetAllProductsAdmin(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAllAdmin(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get admin products: %w", err)
	}
	return s.toProductResponses(products), nil
}

// GetAllProductsAdminPaginated retrieves all products with pagination and optional search for admin
func (s *Service) GetAllProductsAdminPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error) {
	page.Normalize()
	products, total, err := s.productRepo.GetAllAdminPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get admin products: %w", err)
	}
	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return s.toProductResponses(products), pagination, nil
}

// GetProductByID retrieves a single product by ID
func (s *Service) GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get product: %w", err)
	}
	resp := s.toProductResponse(*product)
	return &resp, nil
}

// GetAllCategories retrieves all categories
func (s *Service) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get categories: %w", err)
	}

	responses := make([]dto.CategoryResponse, 0, len(categories))
	for _, c := range categories {
		responses = append(responses, dto.CategoryResponse{
			ID:          c.ID,
			Name:        c.Name,
			Slug:        c.Slug,
			Description: c.Description,
		})
	}
	return responses, nil
}

func (s *Service) toProductResponses(products []model.Product) []dto.ProductResponse {
	responses := make([]dto.ProductResponse, 0, len(products))
	for _, p := range products {
		responses = append(responses, s.toProductResponse(p))
	}
	return responses
}

func (s *Service) toProductResponse(p model.Product) dto.ProductResponse {
	resp := dto.ProductResponse{
		ID:            p.ID,
		ShopID:        p.ShopID,
		Name:          p.Name,
		Slug:          p.Slug,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		ImageURL:      p.ImageURL,
		IsActive:      p.IsActive,
	}

	if p.Shop.ID != "" {
		resp.ShopName = p.Shop.Name
	}
	if p.CategoryID != nil {
		resp.CategoryID = *p.CategoryID
	}
	if p.Category != nil {
		resp.CategoryName = p.Category.Name
	}

	return resp
}
