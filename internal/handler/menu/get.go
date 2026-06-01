package menu

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllMenus handles GET /api/admin/menus
func (h *Handler) GetAllMenus(c *fiber.Ctx) error {
	ctx := c.Context()

	menus, err := h.menuService.GetAllMenus(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve menus",
		})
	}

	return c.JSON(fiber.Map{
		"data": menus,
	})
}
