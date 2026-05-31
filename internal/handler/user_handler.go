package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// UserHandler handles HTTP requests for user management
type UserHandler struct {
	userService *service.UserService
	roleService *service.RoleService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *service.UserService, roleService *service.RoleService) *UserHandler {
	return &UserHandler{
		userService: userService,
		roleService: roleService,
	}
}

// GetAllUsers handles GET /api/admin/users
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
	}
	search := c.Query("search", "")

	users, pagination, err := h.userService.GetAllPaginated(ctx, page, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(fiber.Map{
		"data":       users,
		"pagination": pagination,
	})
}

// CreateUser handles POST /api/admin/users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.CreateUserRequest
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

	user, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": user,
	})
}

// UpdateUser handles PUT /api/admin/users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.userService.UpdateUser(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated",
	})
}

// --- Page Handlers ---

// RenderAdminUsersPage renders the admin users page
func (h *UserHandler) RenderAdminUsersPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
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
