package order

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderAdminOrdersPage renders the admin orders management page
func (h *Handler) RenderAdminOrdersPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	orders, pagination, err := h.orderService.GetAllOrdersPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load orders",
		})
	}

	return c.Render("admin/orders", fiber.Map{
		"Title":      "Manage Orders",
		"Orders":     orders,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
