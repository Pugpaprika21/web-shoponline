package service

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/model"
	"web-shoponline/internal/repository"
)

// ProductService handles business logic for products
type ProductService struct {
	productRepo  *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
}

// NewProductService creates a new ProductService
func NewProductService(productRepo *repository.ProductRepository, categoryRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// GetAllProducts retrieves all active products as DTOs
func (s *ProductService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get products: %w", err)
	}
	return s.toProductResponses(products), nil
}

// GetAllProductsPaginated retrieves active products with pagination and optional search
func (s *ProductService) GetAllProductsPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error) {
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
func (s *ProductService) GetAllProductsAdmin(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.productRepo.GetAllAdmin(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get admin products: %w", err)
	}
	return s.toProductResponses(products), nil
}

// GetAllProductsAdminPaginated retrieves all products with pagination and optional search for admin
func (s *ProductService) GetAllProductsAdminPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error) {
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
func (s *ProductService) GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get product: %w", err)
	}
	resp := s.toProductResponse(*product)
	return &resp, nil
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	slug := repository.GenerateSlug(req.Name)

	product := &model.Product{
		Name:          req.Name,
		Slug:          slug,
		Description:   req.Description,
		Price:         req.Price,
		StockQuantity: req.Stock,
		ImageURL:      req.ImageURL,
		IsActive:      req.IsActive,
	}

	if req.CategoryID != "" {
		product.CategoryID = &req.CategoryID
	}

	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("service: failed to create product: %w", err)
	}

	resp := s.toProductResponse(*product)
	return &resp, nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: product not found: %w", err)
	}

	if req.Name != "" {
		product.Name = req.Name
		product.Slug = repository.GenerateSlug(req.Name)
	}
	if req.CategoryID != "" {
		product.CategoryID = &req.CategoryID
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock > 0 {
		product.StockQuantity = req.Stock
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := s.productRepo.Update(ctx, product); err != nil {
		return nil, fmt.Errorf("service: failed to update product: %w", err)
	}

	resp := s.toProductResponse(*product)
	return &resp, nil
}

// DeleteProduct removes a product
func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.Delete(ctx, id)
}

// GetAllCategories retrieves all categories
func (s *ProductService) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
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

func (s *ProductService) toProductResponses(products []model.Product) []dto.ProductResponse {
	responses := make([]dto.ProductResponse, 0, len(products))
	for _, p := range products {
		responses = append(responses, s.toProductResponse(p))
	}
	return responses
}

func (s *ProductService) toProductResponse(p model.Product) dto.ProductResponse {
	resp := dto.ProductResponse{
		ID:            p.ID,
		Name:          p.Name,
		Slug:          p.Slug,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		ImageURL:      p.ImageURL,
		IsActive:      p.IsActive,
	}

	if p.CategoryID != nil {
		resp.CategoryID = *p.CategoryID
	}
	if p.Category != nil {
		resp.CategoryName = p.Category.Name
	}

	return resp
}
