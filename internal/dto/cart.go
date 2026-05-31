package dto

// AddToCartRequest is the DTO for adding an item to cart
type AddToCartRequest struct {
	ProductID string `json:"product_id" form:"product_id" validate:"required,uuid"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required,gt=0"`
}

// CartResponse is the DTO for returning cart data
type CartResponse struct {
	ID         string             `json:"id"`
	Items      []CartItemResponse `json:"items"`
	TotalItems int                `json:"total_items"`
	TotalPrice float64            `json:"total_price"`
}

// CartItemResponse is the DTO for returning cart item data
type CartItemResponse struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url,omitempty"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}
