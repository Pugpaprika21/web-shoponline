package user

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/handler/common"
)

// GetAllUsers handles GET /api/admin/users
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Context()

	page := dto.PaginationRequest{
		Page:     common.IntQuery(c, "page", 1),
		PageSize: common.IntQuery(c, "page_size", 10),
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
