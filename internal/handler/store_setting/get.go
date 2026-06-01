package storesetting

import (
	"github.com/gofiber/fiber/v2"
)

// GetPublicStoreInfo handles GET /api/store-info (public, no auth)
func (h *Handler) GetPublicStoreInfo(c *fiber.Ctx) error {
	ctx := c.Context()

	settings, err := h.storeSettingService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve store info",
		})
	}

	return c.JSON(fiber.Map{
		"data": settings,
	})
}

// GetStoreSettings handles GET /api/admin/settings/store
func (h *Handler) GetStoreSettings(c *fiber.Ctx) error {
	ctx := c.Context()

	settings, err := h.storeSettingService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve store settings",
		})
	}

	return c.JSON(fiber.Map{
		"data": settings,
	})
}
