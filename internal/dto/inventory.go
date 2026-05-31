package dto

import "time"

// CreateInventoryRequest is the DTO for adding inventory
type CreateInventoryRequest struct {
	ProductID string `json:"product_id" form:"product_id" validate:"required,uuid"`
	Type      string `json:"type" form:"type" validate:"required,oneof=in out adjustment"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required,ne=0"`
	Note      string `json:"note" form:"note" validate:"omitempty,max=500"`
}

// InventoryResponse is the DTO for returning inventory data
type InventoryResponse struct {
	ID          string    `json:"id"`
	ProductID   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Type        string    `json:"type"`
	Quantity    int       `json:"quantity"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
}
