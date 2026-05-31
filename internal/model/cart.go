package model

import "time"

// Cart represents a user's shopping cart
type Cart struct {
	ID        string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    string     `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TableName specifies the table name for Cart
func (Cart) TableName() string {
	return "carts"
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CartID    string    `gorm:"type:uuid;not null" json:"cart_id"`
	ProductID string    `gorm:"type:uuid;not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for CartItem
func (CartItem) TableName() string {
	return "cart_items"
}
