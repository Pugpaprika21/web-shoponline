package inventory

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderAdminInventoryPage renders the admin inventory management page
func (h *Handler) RenderAdminInventoryPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	records, pagination, err := h.inventoryService.GetAllPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load inventory",
		})
	}

	products, err := h.productService.GetAllProductsAdmin(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load products",
		})
	}

	return c.Render("admin/inventory", fiber.Map{
		"Title":      "Manage Inventory",
		"Records":    records,
		"Products":   products,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
