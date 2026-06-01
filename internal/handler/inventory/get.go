package inventory

import (
	"github.com/gofiber/fiber/v2"
)

// GetAll handles GET /api/admin/inventory
func (h *Handler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	records, err := h.inventoryService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve inventory records",
		})
	}

	return c.JSON(fiber.Map{
		"data": records,
	})
}
