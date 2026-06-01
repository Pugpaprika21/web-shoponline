package auth

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
	"web-shoponline/internal/middleware"
)

// Register handles POST /api/auth/register
func (h *Handler) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.RegisterRequest
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

	user, err := h.authService.Register(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate JWT token for auto-login after register
	token, err := middleware.GenerateToken(h.jwtConfig, user.ID, user.Email, user.RoleNames)
	if err != nil {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":    user,
			"message": "Registration successful",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   h.jwtConfig.ExpireHour * 3600,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    user,
		"token":   token,
		"message": "Registration successful",
	})
}
