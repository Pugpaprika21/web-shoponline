package product

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteProduct handles DELETE /api/admin/products/:id
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.productService.DeleteProduct(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
