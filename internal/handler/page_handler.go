package handler

import (
	"github.com/gofiber/fiber/v2"
)

// PageHandler handles static page rendering
type PageHandler struct{}

// NewPageHandler creates a new PageHandler
func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

// RenderHomePage renders the home page
func (h *PageHandler) RenderHomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Welcome to ShopOnline",
	}, "layouts/main")
}

// RenderAdminDashboard renders the admin dashboard
func (h *PageHandler) RenderAdminDashboard(c *fiber.Ctx) error {
	return c.Render("admin/dashboard", fiber.Map{
		"Title": "Admin Dashboard",
	}, "layouts/admin")
}
