package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	productService *service.ProductService
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetAllProducts handles GET /api/products
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	ctx := c.Context()

	products, err := h.productService.GetAllProducts(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

// GetProductByID handles GET /api/products/:id
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": product,
	})
}

// CreateProduct handles POST /api/admin/products
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": formatValidationErrors(err),
		})
	}

	product, err := h.productService.CreateProduct(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": product,
	})
}

// UpdateProduct handles PUT /api/admin/products/:id
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": formatValidationErrors(err),
		})
	}

	product, err := h.productService.UpdateProduct(ctx, id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"data": product,
	})
}

// DeleteProduct handles DELETE /api/admin/products/:id
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.productService.DeleteProduct(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// --- HTMX Page Handlers ---

// RenderProductsPage renders the products page for storefront
func (h *ProductHandler) RenderProductsPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 12),
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

// RenderAdminProductsPage renders the admin products management page
func (h *ProductHandler) RenderAdminProductsPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
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

// intQuery parses an integer query parameter with a default value
func intQuery(c *fiber.Ctx, key string, defaultVal int) int {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return n
}
