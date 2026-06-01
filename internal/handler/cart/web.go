package cart

import (
	"github.com/gofiber/fiber/v2"
)

// RenderCartPage renders the shopping cart page
func (h *Handler) RenderCartPage(c *fiber.Ctx) error {
	return c.Render("cart", fiber.Map{
		"Title": "Shopping Cart",
	}, "layouts/main")
}
