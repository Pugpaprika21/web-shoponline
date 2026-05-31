package handler

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/service"
)

// CartHandler handles HTTP requests for shopping cart
type CartHandler struct {
	cartService *service.CartService
}

// NewCartHandler creates a new CartHandler
func NewCartHandler(cartService *service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// getUserID extracts user ID from JWT locals (set by middleware)
func getUserID(c *fiber.Ctx) string {
	if id, ok := c.Locals("userID").(string); ok {
		return id
	}
	return ""
}

// GetCart handles GET /api/cart
func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := getUserID(c)

	cart, err := h.cartService.GetCart(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get cart",
		})
	}

	return c.JSON(fiber.Map{
		"data": cart,
	})
}

// AddToCart handles POST /api/cart/items
func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := getUserID(c)

	var req dto.AddToCartRequest
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

	if err := h.cartService.AddToCart(ctx, userID, req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Item added to cart",
	})
}

// RemoveFromCart handles DELETE /api/cart/items/:id
func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := getUserID(c)
	itemID := c.Params("id")

	if err := h.cartService.RemoveFromCart(ctx, userID, itemID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove item",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item removed from cart",
	})
}

// GetCartCount handles GET /api/cart/count
func (h *CartHandler) GetCartCount(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := getUserID(c)

	count, err := h.cartService.GetCartItemCount(ctx, userID)
	if err != nil {
		return c.JSON(fiber.Map{"count": 0})
	}

	return c.JSON(fiber.Map{"count": count})
}

// --- Page Handlers ---

// RenderCartPage renders the shopping cart page
func (h *CartHandler) RenderCartPage(c *fiber.Ctx) error {
	return c.Render("cart", fiber.Map{
		"Title": "Shopping Cart",
	}, "layouts/main")
}
