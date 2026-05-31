package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RequireAuth checks if user is authenticated (legacy cookie-based, kept for compatibility)
func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Cookies("user_id")
		if userID == "" {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Authentication required",
				})
			}
			return c.Redirect("/login")
		}
		c.Locals("userID", userID)
		c.Locals("userRole", c.Cookies("user_role"))
		return c.Next()
	}
}

// RequireAdmin checks if user has admin role (legacy cookie-based, kept for compatibility)
func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Cookies("user_id")
		userRole := c.Cookies("user_role")

		if userID == "" || userRole != "admin" {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Admin access required",
				})
			}
			return c.Redirect("/admin/login")
		}

		c.Locals("userID", userID)
		c.Locals("userRole", userRole)
		return c.Next()
	}
}

func isAPIRequest(c *fiber.Ctx) bool {
	accept := c.Get("Accept")
	contentType := c.Get("Content-Type")
	path := c.Path()

	return strings.Contains(accept, "application/json") ||
		strings.Contains(contentType, "application/json") ||
		(len(path) > 4 && path[:4] == "/api")
}
