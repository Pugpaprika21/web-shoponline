package order

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllOrders handles GET /api/admin/orders
func (h *Handler) GetAllOrders(c *fiber.Ctx) error {
	ctx := c.Context()

	orders, err := h.orderService.GetAllOrders(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}

// GetOrderByID handles GET /api/admin/orders/:id
func (h *Handler) GetOrderByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	order, err := h.orderService.GetOrderByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": order,
	})
}
