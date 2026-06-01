package role

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
)

// UpdateRole handles PUT /api/admin/roles/:id
func (h *Handler) UpdateRole(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.roleService.UpdateRole(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update role",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Role updated",
	})
}
