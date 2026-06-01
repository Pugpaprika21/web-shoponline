package auth

import (
	"github.com/gofiber/fiber/v2"
)

// Logout handles POST /api/auth/logout
func (h *Handler) Logout(c *fiber.Ctx) error {
	// Clear token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
