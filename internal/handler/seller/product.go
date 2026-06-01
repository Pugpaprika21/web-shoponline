package seller

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
	"web-shoponline/internal/model"
)

// GetMyProducts handles GET /api/seller/products
func (h *Handler) GetMyProducts(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found. Please create a shop first.",
		})
	}

	products, err := h.productRepo.GetByShopID(ctx, shop.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

// CreateProduct handles POST /api/seller/products
func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found. Please create a shop first.",
		})
	}

	var req dto.SellerCreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := common.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": common.FormatValidationErrors(err),
		})
	}

	// Generate slug from name
	slug := strings.ToLower(strings.ReplaceAll(req.Name, " ", "-"))

	product := &model.Product{
		ShopID:        shop.ID,
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

	if err := h.productRepo.Create(ctx, product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    product,
		"message": "Product created successfully",
	})
}

// UpdateProduct handles PUT /api/seller/products/:id
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)
	productID := c.Params("id")

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found",
		})
	}

	// Get product and verify ownership
	product, err := h.productRepo.GetByID(ctx, productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if product.ShopID != shop.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You can only edit your own products",
		})
	}

	var req dto.SellerUpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := common.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": common.FormatValidationErrors(err),
		})
	}

	// Update fields
	if req.Name != "" {
		product.Name = req.Name
		product.Slug = strings.ToLower(strings.ReplaceAll(req.Name, " ", "-"))
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.StockQuantity = req.Stock
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}
	if req.CategoryID != "" {
		product.CategoryID = &req.CategoryID
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := h.productRepo.Update(ctx, product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"data":    product,
		"message": "Product updated successfully",
	})
}

// DeleteProduct handles DELETE /api/seller/products/:id
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)
	productID := c.Params("id")

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found",
		})
	}

	// Get product and verify ownership
	product, err := h.productRepo.GetByID(ctx, productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if product.ShopID != shop.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You can only delete your own products",
		})
	}

	if err := h.productRepo.Delete(ctx, productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
