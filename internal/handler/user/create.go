package user

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// CreateUser handles POST /api/admin/users
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateUserRequest
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

	u, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": u,
	})
}
