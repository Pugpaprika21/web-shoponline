package product

import "web-shoponline/internal/service"

// Handler handles HTTP requests for products
type Handler struct {
	productService service.ProductServiceInterface
}

// NewHandler creates a new product Handler
func NewHandler(productService service.ProductServiceInterface) *Handler {
	return &Handler{productService: productService}
}
