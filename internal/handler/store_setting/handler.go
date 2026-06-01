package storesetting

import "web-shoponline/internal/service"

// Handler handles HTTP requests for store settings
type Handler struct {
	storeSettingService service.StoreSettingServiceInterface
}

// NewHandler creates a new store setting Handler
func NewHandler(storeSettingService service.StoreSettingServiceInterface) *Handler {
	return &Handler{storeSettingService: storeSettingService}
}
