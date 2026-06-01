package auth

import (
	"web-shoponline/internal/config"
	"web-shoponline/internal/service"
)

// Handler handles authentication HTTP requests
type Handler struct {
	authService service.AuthServiceInterface
	jwtConfig   config.JWTConfig
}

// NewHandler creates a new auth Handler
func NewHandler(authService service.AuthServiceInterface, jwtConfig config.JWTConfig) *Handler {
	return &Handler{
		authService: authService,
		jwtConfig:   jwtConfig,
	}
}
