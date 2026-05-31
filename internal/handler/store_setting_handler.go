package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// StoreSettingHandler handles HTTP requests for store settings
type StoreSettingHandler struct {
	storeSettingService *service.StoreSettingService
}

// NewStoreSettingHandler creates a new StoreSettingHandler
func NewStoreSettingHandler(storeSettingService *service.StoreSettingService) *StoreSettingHandler {
	return &StoreSettingHandler{storeSettingService: storeSettingService}
}

// GetStoreSettings handles GET /api/admin/settings/store
func (h *StoreSettingHandler) GetStoreSettings(c *fiber.Ctx) error {
	ctx := c.Context()

	settings, err := h.storeSettingService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve store settings",
		})
	}

	return c.JSON(fiber.Map{
		"data": settings,
	})
}

// UpdateStoreSettings handles PUT /api/admin/settings/store
func (h *StoreSettingHandler) UpdateStoreSettings(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.StoreSettingsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.storeSettingService.UpdateAll(ctx, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update store settings",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Store settings updated",
	})
}

// --- Page Handlers ---

// RenderAdminStoreSettingsPage renders the admin store settings page
func (h *StoreSettingHandler) RenderAdminStoreSettingsPage(c *fiber.Ctx) error {
	ctx := c.Context()

	settings, err := h.storeSettingService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load store settings",
		})
	}

	return c.Render("admin/settings_store", fiber.Map{
		"Title":    "Store Settings",
		"Settings": settings,
	}, "layouts/admin")
}
