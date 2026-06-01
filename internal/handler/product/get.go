package product

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts handles GET /api/products
func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	ctx := c.Context()

	products, err := h.productService.GetAllProducts(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

// GetProductByID handles GET /api/products/:id
func (h *Handler) GetProductByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": product,
	})
}
