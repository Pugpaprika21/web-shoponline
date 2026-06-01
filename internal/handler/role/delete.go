package role

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteRole handles DELETE /api/admin/roles/:id
func (h *Handler) DeleteRole(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.roleService.DeleteRole(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete role",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
