package inventory

import "web-shoponline/internal/service"

// Handler handles HTTP requests for inventory management
type Handler struct {
	inventoryService service.InventoryServiceInterface
	productService   service.ProductServiceInterface
}

// NewHandler creates a new inventory Handler
func NewHandler(inventoryService service.InventoryServiceInterface, productService service.ProductServiceInterface) *Handler {
	return &Handler{
		inventoryService: inventoryService,
		productService:   productService,
	}
}
