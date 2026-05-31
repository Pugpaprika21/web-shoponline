package dto

import "time"

// OrderResponse is the DTO for returning order data
type OrderResponse struct {
	ID              string              `json:"id"`
	UserEmail       string              `json:"user_email,omitempty"`
	Status          string              `json:"status"`
	TotalAmount     float64             `json:"total_amount"`
	ShippingAddress string              `json:"shipping_address,omitempty"`
	Items           []OrderItemResponse `json:"items,omitempty"`
	CreatedAt       time.Time           `json:"created_at"`
}

// OrderItemResponse is the DTO for returning order item data
type OrderItemResponse struct {
	ID          string  `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Subtotal    float64 `json:"subtotal"`
}

// CreateOrderRequest is the DTO for creating a new order
type CreateOrderRequest struct {
	ShippingAddress string             `json:"shipping_address" form:"shipping_address" validate:"required,min=10,max=500"`
	Items           []OrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

// OrderItemRequest is the DTO for an item in a new order
type OrderItemRequest struct {
	ProductID string `json:"product_id" validate:"required,uuid"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
}

// UpdateOrderStatusRequest is the DTO for updating order status
type UpdateOrderStatusRequest struct {
	Status string `json:"status" form:"status" validate:"required,oneof=pending confirmed shipped delivered cancelled"`
}
