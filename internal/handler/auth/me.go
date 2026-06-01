package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"web-shoponline/internal/middleware"
)

// GetCurrentUser handles GET /api/auth/me
func (h *Handler) GetCurrentUser(c *fiber.Ctx) error {
	ctx := c.Context()

	// Try to get user from JWT (via locals set by middleware or manual extraction)
	tokenString := ""
	authHeader := c.Get("Authorization")
	if authHeader != "" {
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}
	}
	if tokenString == "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authenticated",
		})
	}

	// Validate token manually here since this endpoint doesn't use middleware
	claims, err := parseToken(h.jwtConfig.Secret, tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	user, err := h.authService.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// parseToken validates a JWT token and returns claims
func parseToken(secret, tokenString string) (*middleware.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &middleware.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*middleware.JWTClaims)
	if !ok || !token.Valid {
		return nil, fiber.ErrUnauthorized
	}
	return claims, nil
}
