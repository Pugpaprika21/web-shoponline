package auth

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
	"web-shoponline/internal/middleware"
)

// Login handles POST /api/auth/login
func (h *Handler) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := common.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": common.FormatValidationErrors(err),
		})
	}

	user, err := h.authService.Login(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate JWT token with roles array
	token, err := middleware.GenerateToken(h.jwtConfig, user.ID, user.Email, user.RoleNames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set token in HttpOnly cookie (for browser-based requests)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   h.jwtConfig.ExpireHour * 3600,
	})

	return c.JSON(fiber.Map{
		"data":    user,
		"token":   token,
		"message": "Login successful",
	})
}
