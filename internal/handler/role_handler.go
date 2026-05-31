package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// RoleHandler handles HTTP requests for role management
type RoleHandler struct {
	roleService *service.RoleService
}

// NewRoleHandler creates a new RoleHandler
func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// GetAllRoles handles GET /api/admin/roles
func (h *RoleHandler) GetAllRoles(c *fiber.Ctx) error {
	ctx := c.Context()

	roles, err := h.roleService.GetAllRoles(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve roles",
		})
	}

	return c.JSON(fiber.Map{
		"data": roles,
	})
}

// GetAllPermissions handles GET /api/admin/permissions
func (h *RoleHandler) GetAllPermissions(c *fiber.Ctx) error {
	ctx := c.Context()

	permissions, err := h.roleService.GetAllPermissions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve permissions",
		})
	}

	return c.JSON(fiber.Map{
		"data": permissions,
	})
}

// CreateRole handles POST /api/admin/roles
func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateRoleRequest
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

	role, err := h.roleService.CreateRole(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create role",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": role,
	})
}

// UpdateRole handles PUT /api/admin/roles/:id
func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.roleService.UpdateRole(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update role",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Role updated",
	})
}

// DeleteRole handles DELETE /api/admin/roles/:id
func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if err := h.roleService.DeleteRole(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete role",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// --- Page Handlers ---

// RenderAdminRolesPage renders the admin roles settings page
func (h *RoleHandler) RenderAdminRolesPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
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
