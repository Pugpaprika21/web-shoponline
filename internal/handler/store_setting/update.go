package storesetting

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
)

// UpdateStoreSettings handles PUT /api/admin/settings/store
func (h *Handler) UpdateStoreSettings(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.StoreSettingsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.storeSettingService.UpdateAll(ctx, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update store settings",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Store settings updated",
	})
}
