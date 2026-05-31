package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// InventoryHandler handles HTTP requests for inventory management
type InventoryHandler struct {
	inventoryService *service.InventoryService
	productService   *service.ProductService
}

// NewInventoryHandler creates a new InventoryHandler
func NewInventoryHandler(inventoryService *service.InventoryService, productService *service.ProductService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
		productService:   productService,
	}
}

// GetAll handles GET /api/admin/inventory
func (h *InventoryHandler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	records, err := h.inventoryService.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve inventory records",
		})
	}

	return c.JSON(fiber.Map{
		"data": records,
	})
}

// AddInventory handles POST /api/admin/inventory
func (h *InventoryHandler) AddInventory(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := c.Locals("userID").(string)

	var req dto.CreateInventoryRequest
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

	record, err := h.inventoryService.AddInventory(ctx, req, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": record,
	})
}

// --- Page Handlers ---

// RenderAdminInventoryPage renders the admin inventory management page
func (h *InventoryHandler) RenderAdminInventoryPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
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
