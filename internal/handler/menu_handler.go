package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// MenuHandler handles HTTP requests for menu management
type MenuHandler struct {
	menuService *service.MenuService
}

// NewMenuHandler creates a new MenuHandler
func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

// GetAllMenus handles GET /api/admin/menus
func (h *MenuHandler) GetAllMenus(c *fiber.Ctx) error {
	ctx := c.Context()

	menus, err := h.menuService.GetAllMenus(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve menus",
		})
	}

	return c.JSON(fiber.Map{
		"data": menus,
	})
}

// CreateMenu handles POST /api/admin/menus
func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateMenuRequest
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

	menu, err := h.menuService.CreateMenu(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": menu,
	})
}

// UpdateMenu handles PUT /api/admin/menus/:id
func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateMenuRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.menuService.UpdateMenu(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Menu updated",
	})
}

// DeleteMenu handles DELETE /api/admin/menus/:id
func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.menuService.DeleteMenu(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// --- Page Handlers ---

// RenderAdminMenusPage renders the admin menus settings page
func (h *MenuHandler) RenderAdminMenusPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	menus, pagination, err := h.menuService.GetAllMenusPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"message": "Failed to load menus",
		})
	}

	return c.Render("admin/settings_menus", fiber.Map{
		"Title":      "Menu Settings",
		"Menus":      menus,
		"Pagination": pagination,
		"Search":     search,
	}, "layouts/admin")
}
