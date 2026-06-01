package page

import (
	"github.com/gofiber/fiber/v2"
)

// RenderHomePage renders the home page
func (h *Handler) RenderHomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Welcome to ShopOnline",
	}, "layouts/main")
}

// RenderAdminDashboard renders the admin dashboard
func (h *Handler) RenderAdminDashboard(c *fiber.Ctx) error {
	return c.Render("admin/dashboard", fiber.Map{
		"Title": "Admin Dashboard",
	}, "layouts/admin")
}
