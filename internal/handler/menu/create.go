package menu

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// CreateMenu handles POST /api/admin/menus
func (h *Handler) CreateMenu(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateMenuRequest
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

	m, err := h.menuService.CreateMenu(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": m,
	})
}
