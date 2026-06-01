package role

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllRoles handles GET /api/admin/roles
func (h *Handler) GetAllRoles(c *fiber.Ctx) error {
	ctx := c.Context()

	roles, err := h.roleService.GetAllRoles(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve roles",
		})
	}

	return c.JSON(fiber.Map{
		"data": roles,
	})
}

// GetAllPermissions handles GET /api/admin/permissions
func (h *Handler) GetAllPermissions(c *fiber.Ctx) error {
	ctx := c.Context()

	permissions, err := h.roleService.GetAllPermissions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve permissions",
		})
	}

	return c.JSON(fiber.Map{
		"data": permissions,
	})
}
