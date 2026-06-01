package cart

import "web-shoponline/internal/service"

// Handler handles HTTP requests for shopping cart
type Handler struct {
	cartService service.CartServiceInterface
}

// NewHandler creates a new cart Handler
func NewHandler(cartService service.CartServiceInterface) *Handler {
	return &Handler{cartService: cartService}
}
