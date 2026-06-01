package role

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// CreateRole handles POST /api/admin/roles
func (h *Handler) CreateRole(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateRoleRequest
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

	rl, err := h.roleService.CreateRole(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create role",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": rl,
	})
}
