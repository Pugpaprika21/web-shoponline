package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// OrderHandler handles HTTP requests for orders
type OrderHandler struct {
	orderService *service.OrderService
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// GetAllOrders handles GET /api/admin/orders
func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	ctx := c.Context()

	orders, err := h.orderService.GetAllOrders(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}

// GetOrderByID handles GET /api/admin/orders/:id
func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	order, err := h.orderService.GetOrderByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": order,
	})
}

// UpdateOrderStatus handles PATCH /api/admin/orders/:id/status
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var req dto.UpdateOrderStatusRequest
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

	if err := h.orderService.UpdateOrderStatus(ctx, id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order status updated",
	})
}

// --- HTMX Page Handlers ---

// RenderAdminOrdersPage renders the admin orders management page
func (h *OrderHandler) RenderAdminOrdersPage(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     intQuery(c, "page", 1),
		PageSize: intQuery(c, "page_size", 10),
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
