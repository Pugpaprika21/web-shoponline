package auth

import (
	"github.com/gofiber/fiber/v2"
)

// RenderLoginPage renders the user login page
func (h *Handler) RenderLoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	}, "layouts/main")
}

// RenderRegisterPage renders the user registration page
func (h *Handler) RenderRegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	}, "layouts/main")
}

// RenderAdminLoginPage renders the admin login page
func (h *Handler) RenderAdminLoginPage(c *fiber.Ctx) error {
	return c.Render("admin/login", fiber.Map{
		"Title": "Admin Login",
	})
}
