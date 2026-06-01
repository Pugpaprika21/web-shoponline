package order

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// UpdateOrderStatus handles PATCH /api/admin/orders/:id/status
func (h *Handler) UpdateOrderStatus(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateOrderStatusRequest
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

	if err := h.orderService.UpdateOrderStatus(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order status updated",
	})
}
