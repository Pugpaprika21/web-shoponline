package seller

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
	"web-shoponline/internal/model"
)

// GetMyShop handles GET /api/seller/shop
func (h *Handler) GetMyShop(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found. Please create a shop first.",
		})
	}

	return c.JSON(fiber.Map{
		"data": shop,
	})
}

// UpdateMyShop handles PUT /api/seller/shop
func (h *Handler) UpdateMyShop(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found",
		})
	}

	var req dto.UpdateShopRequest
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

	shop.Name = req.Name
	shop.Description = req.Description
	shop.LogoURL = req.LogoURL
	shop.BannerURL = req.BannerURL
	shop.Phone = req.Phone
	shop.Email = req.Email
	shop.Address = req.Address

	if err := h.shopRepo.Update(ctx, shop); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update shop",
		})
	}

	return c.JSON(fiber.Map{
		"data":    shop,
		"message": "Shop updated successfully",
	})
}

// CreateMyShop handles POST /api/seller/shop
func (h *Handler) CreateMyShop(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	// Check if seller already has a shop
	existing, _ := h.shopRepo.GetByOwnerID(ctx, userID)
	if existing != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "You already have a shop",
		})
	}

	var req dto.CreateShopRequest
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

	shop := &model.Shop{
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,
		OwnerID:     userID,
		IsActive:    true,
	}

	if err := h.shopRepo.Create(ctx, shop); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create shop",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    shop,
		"message": "Shop created successfully",
	})
}
