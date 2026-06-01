package seller

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/handler/common"
)

// RenderSellerDashboard renders the seller dashboard overview page
func (h *Handler) RenderSellerDashboard(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, _ := h.shopRepo.GetByOwnerID(ctx, userID)

	var productCount int
	var orderCount int
	if shop != nil {
		products, _ := h.productRepo.GetByShopID(ctx, shop.ID)
		productCount = len(products)
		// Count orders that contain products from this shop
		orders, _ := h.orderRepo.GetByShopID(ctx, shop.ID)
		orderCount = len(orders)
	}

	return c.Render("seller/dashboard", fiber.Map{
		"Title":        "Seller Dashboard",
		"Shop":         shop,
		"ProductCount": productCount,
		"OrderCount":   orderCount,
	}, "layouts/seller")
}

// RenderSellerShopSettings renders the shop settings page for the seller
func (h *Handler) RenderSellerShopSettings(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, _ := h.shopRepo.GetByOwnerID(ctx, userID)

	return c.Render("seller/shop_settings", fiber.Map{
		"Title": "Shop Settings",
		"Shop":  shop,
	}, "layouts/seller")
}

// RenderSellerProducts renders the products management page for the seller
func (h *Handler) RenderSellerProducts(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, _ := h.shopRepo.GetByOwnerID(ctx, userID)

	return c.Render("seller/products", fiber.Map{
		"Title": "My Products",
		"Shop":  shop,
	}, "layouts/seller")
}

// RenderSellerOrders renders the orders page for the seller
func (h *Handler) RenderSellerOrders(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := common.GetUserID(c)

	shop, _ := h.shopRepo.GetByOwnerID(ctx, userID)

	return c.Render("seller/orders", fiber.Map{
		"Title": "My Orders",
		"Shop":  shop,
	}, "layouts/seller")
}
