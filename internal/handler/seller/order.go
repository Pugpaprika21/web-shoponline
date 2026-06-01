package seller

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/handler/common"
)

// GetMyOrders handles GET /api/seller/orders
func (h *Handler) GetMyOrders(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, err := h.shopRepo.GetByOwnerID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shop not found. Please create a shop first.",
		})
	}

	orders, err := h.orderRepo.GetByShopID(ctx, shop.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}
