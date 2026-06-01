package shop

import (
	"github.com/gofiber/fiber/v2"
)

// RenderShopPage renders a shop's storefront page
func (h *Handler) RenderShopPage(c *fiber.Ctx) error {
	ctx := c.Context()
	slugOrID := c.Params("slug")

	// Try by slug first, then by ID
	shop, err := h.shopRepo.GetBySlug(ctx, slugOrID)
	if err != nil {
		shop, err = h.shopRepo.GetByID(ctx, slugOrID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
				"message": "Shop not found",
			})
		}
	}

	products, err := h.productRepo.GetByShopID(ctx, shop.ID)
	if err != nil {
		products = nil
	}

	return c.Render("shop", fiber.Map{
		"Title":    shop.Name,
		"Shop":     shop,
		"Products": products,
	}, "layouts/main")
}
