package menu

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteMenu handles DELETE /api/admin/menus/:id
func (h *Handler) DeleteMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.menuService.DeleteMenu(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
