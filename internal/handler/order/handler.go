package order

import "web-shoponline/internal/service"

// Handler handles HTTP requests for orders
type Handler struct {
	orderService service.OrderServiceInterface
}

// NewHandler creates a new order Handler
func NewHandler(orderService service.OrderServiceInterface) *Handler {
	return &Handler{orderService: orderService}
}
