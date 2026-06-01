package role

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderAdminRolesPage renders the admin roles settings page
func (h *Handler) RenderAdminRolesPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	roles, pagination, err := h.roleService.GetAllRolesPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load roles",
		})
	}

	permissions, err := h.roleService.GetAllPermissions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load permissions",
		})
	}

	return c.Render("admin/settings_roles", fiber.Map{
		"Title":       "Role Settings",
		"Roles":       roles,
		"Permissions": permissions,
		"Pagination":  pagination,
		"Search":      search,
	}, "layouts/admin")
}
