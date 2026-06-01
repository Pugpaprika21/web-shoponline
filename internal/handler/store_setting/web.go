package storesetting

import (
	"github.com/gofiber/fiber/v2"
)

// RenderAdminStoreSettingsPage renders the admin store settings page
func (h *Handler) RenderAdminStoreSettingsPage(c *fiber.Ctx) error {
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
