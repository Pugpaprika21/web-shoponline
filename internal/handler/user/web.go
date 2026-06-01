package user

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderAdminUsersPage renders the admin users page
func (h *Handler) RenderAdminUsersPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	users, pagination, err := h.userService.GetAllPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load users",
		})
	}

	roles, err := h.roleService.GetAllRoles(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load roles",
		})
	}

	return c.Render("admin/users", fiber.Map{
		"Title":      "User Management",
		"Users":      users,
		"Roles":      roles,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
