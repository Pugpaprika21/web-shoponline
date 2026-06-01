package cart

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/handler/common"
)

// GetCart handles GET /api/cart
func (h *Handler) GetCart(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	cart, err := h.cartService.GetCart(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get cart",
		})
	}

	return c.JSON(fiber.Map{
		"data": cart,
	})
}

// GetCartCount handles GET /api/cart/count
func (h *Handler) GetCartCount(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	count, err := h.cartService.GetCartItemCount(ctx, userID)
	if err != nil {
		return c.JSON(fiber.Map{"count": 0})
	}

	return c.JSON(fiber.Map{"count": count})
}
