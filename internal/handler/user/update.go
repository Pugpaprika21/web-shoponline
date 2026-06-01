package user

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
)

// UpdateUser handles PUT /api/admin/users/:id
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.userService.UpdateUser(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated",
	})
}
