package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"web-shoponline/internal/config"
)

// JWTClaims holds the JWT token claims
type JWTClaims struct {
	UserID string   `json:"user_id"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(cfg config.JWTConfig, userID, email string, roles []string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.ExpireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// RequireJWT validates the JWT token from Authorization header or cookie
func RequireJWT(cfg config.JWTConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := extractToken(c)
		if tokenString == "" {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Authentication required",
				})
			}
			return c.Redirect("/login")
		}

		claims, err := validateToken(cfg.Secret, tokenString)
		if err != nil {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid or expired token",
				})
			}
			return c.Redirect("/login")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRoles", claims.Roles)
		return c.Next()
	}
}

// RequireAdminJWT validates JWT and checks admin role
func RequireAdminJWT(cfg config.JWTConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := extractToken(c)
		if tokenString == "" {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Authentication required",
				})
			}
			return c.Redirect("/admin/login")
		}

		claims, err := validateToken(cfg.Secret, tokenString)
		if err != nil {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid or expired token",
				})
			}
			return c.Redirect("/admin/login")
		}

		if !hasRole(claims.Roles, "admin") {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Admin access required",
				})
			}
			return c.Redirect("/admin/login")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRoles", claims.Roles)
		return c.Next()
	}
}

// RequireSellerJWT validates JWT and checks seller or admin role
func RequireSellerJWT(cfg config.JWTConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := extractToken(c)
		if tokenString == "" {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Authentication required",
				})
			}
			return c.Redirect("/login")
		}

		claims, err := validateToken(cfg.Secret, tokenString)
		if err != nil {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid or expired token",
				})
			}
			return c.Redirect("/login")
		}

		if !hasRole(claims.Roles, "seller") && !hasRole(claims.Roles, "admin") {
			if isAPIRequest(c) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Seller access required",
				})
			}
			return c.Redirect("/login")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRoles", claims.Roles)
		return c.Next()
	}
}

// hasRole checks if a role name exists in the roles slice
func hasRole(roles []string, roleName string) bool {
	for _, r := range roles {
		if r == roleName {
			return true
		}
	}
	return false
}

// extractToken gets the token from Authorization header or cookie
func extractToken(c *fiber.Ctx) string {
	// 1. Check Authorization header: Bearer <token>
	authHeader := c.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			return parts[1]
		}
	}

	// 2. Fallback to cookie
	token := c.Cookies("token")
	if token != "" {
		return token
	}

	return ""
}

// validateToken parses and validates a JWT token string
func validateToken(secret, tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
