package menu

import "web-shoponline/internal/service"

// Handler handles HTTP requests for menu management
type Handler struct {
	menuService service.MenuServiceInterface
}

// NewHandler creates a new menu Handler
func NewHandler(menuService service.MenuServiceInterface) *Handler {
	return &Handler{menuService: menuService}
}
