package cart

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/handler/common"
)

// RemoveFromCart handles DELETE /api/cart/items/:id
func (h *Handler) RemoveFromCart(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)
	itemID := c.Params("id")

	if err := h.cartService.RemoveFromCart(ctx, userID, itemID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove item",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item removed from cart",
	})
}
