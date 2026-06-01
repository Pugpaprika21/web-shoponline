package menu

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
)

// UpdateMenu handles PUT /api/admin/menus/:id
func (h *Handler) UpdateMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateMenuRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.menuService.UpdateMenu(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Menu updated",
	})
}
