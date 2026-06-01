package menu

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderAdminMenusPage renders the admin menus settings page
func (h *Handler) RenderAdminMenusPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	menus, pagination, err := h.menuService.GetAllMenusPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load menus",
		})
	}

	// Get all menus (flat) for parent dropdown
	allMenus, _ := h.menuService.GetAllMenus(ctx)

	// Flatten tree for table rendering
	flatMenus := dto.FlattenMenuTree(menus)

	return c.Render("admin/settings_menus", fiber.Map{
		"Title":      "Menu Settings",
		"Menus":      menus,
		"FlatMenus":  flatMenus,
		"AllMenus":   allMenus,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
