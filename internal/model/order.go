package model

import (
	"time"

	"gorm.io/gorm"
)

// Order represents the order database model
type Order struct {
	ID              string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID          string         `gorm:"type:uuid;not null" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status          string         `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	TotalAmount     float64        `gorm:"type:decimal(10,2);not null;default:0" json:"total_amount"`
	ShippingAddress string         `gorm:"type:text" json:"shipping_address"`
	Items           []OrderItem    `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Order
func (Order) TableName() string {
	return "orders"
}

// OrderItem represents an item within an order
type OrderItem struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrderID   string    `gorm:"type:uuid;not null" json:"order_id"`
	ProductID string    `gorm:"type:uuid;not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"not null;default:1" json:"quantity"`
	UnitPrice float64   `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for OrderItem
func (OrderItem) TableName() string {
	return "order_items"
}
