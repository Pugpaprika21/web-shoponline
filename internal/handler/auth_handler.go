package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"web-shoponline/internal/config"
	"web-shoponline/internal/dto"
	"web-shoponline/internal/middleware"
	"web-shoponline/internal/service"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authService *service.AuthService
	jwtConfig   config.JWTConfig
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *service.AuthService, jwtConfig config.JWTConfig) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtConfig:   jwtConfig,
	}
}

// Login handles POST /api/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": formatValidationErrors(err),
		})
	}

	user, err := h.authService.Login(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(h.jwtConfig, user.ID, user.Email, user.RoleName)
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

// Register handles POST /api/auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": formatValidationErrors(err),
		})
	}

	user, err := h.authService.Register(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Generate JWT token for auto-login after register
	token, err := middleware.GenerateToken(h.jwtConfig, user.ID, user.Email, user.RoleName)
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

// Logout handles POST /api/auth/logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
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

// GetCurrentUser handles GET /api/auth/me
func (h *AuthHandler) GetCurrentUser(c *fiber.Ctx) error {
	ctx := c.Context()

	// Try to get user from JWT (via locals set by middleware or manual extraction)
	tokenString := ""
	authHeader := c.Get("Authorization")
	if authHeader != "" {
		parts := make([]byte, 0)
		_ = parts
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

// --- Page Handlers ---

// RenderLoginPage renders the user login page
func (h *AuthHandler) RenderLoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	}, "layouts/main")
}

// RenderRegisterPage renders the user registration page
func (h *AuthHandler) RenderRegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	}, "layouts/main")
}

// RenderAdminLoginPage renders the admin login page
func (h *AuthHandler) RenderAdminLoginPage(c *fiber.Ctx) error {
	return c.Render("admin/login", fiber.Map{
		"Title": "Admin Login",
	})
}
