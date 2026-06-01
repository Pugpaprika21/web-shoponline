package product

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// RenderProductsPage renders the products page for storefront
func (h *Handler) RenderProductsPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 12),
	}
	search := c.Query("search", "")

	products, pagination, err := h.productService.GetAllProductsPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load products",
		})
	}

	return c.Render("products", fiber.Map{
		"Title":      "Products",
		"Products":   products,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/main")
}

// RenderProductDetailPage renders a single product detail page
func (h *Handler) RenderProductDetailPage(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
			"message": "Product not found",
		})
	}

	return c.Render("product_detail", fiber.Map{
		"Title":   product.Name,
		"Product": product,
	}, "layouts/main")
}

// RenderAdminProductsPage renders the admin products management page
func (h *Handler) RenderAdminProductsPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	products, pagination, err := h.productService.GetAllProductsAdminPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load products",
		})
	}

	categories, _ := h.productService.GetAllCategories(ctx)

	return c.Render("admin/products", fiber.Map{
		"Title":      "Manage Products",
		"Products":   products,
		"Categories": categories,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
