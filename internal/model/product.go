package model

import (
	"time"

	"gorm.io/gorm"
)

// Category represents the category database model
type Category struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName specifies the table name for Category
func (Category) TableName() string {
	return "categories"
}

// Product represents the product database model
type Product struct {
	ID            string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CategoryID    *string        `gorm:"type:uuid" json:"category_id"`
	Category      *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Slug          string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description   string         `gorm:"type:text" json:"description"`
	Price         float64        `gorm:"type:decimal(10,2);not null;default:0" json:"price"`
	StockQuantity int            `gorm:"not null;default:0" json:"stock_quantity"`
	ImageURL      string         `gorm:"type:varchar(500)" json:"image_url"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Product
func (Product) TableName() string {
	return "products"
}

// Inventory represents stock movement/adjustment records
type Inventory struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProductID string    `gorm:"type:uuid;not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // in, out, adjustment
	Quantity  int       `gorm:"not null" json:"quantity"`
	Note      string    `gorm:"type:text" json:"note"`
	CreatedBy string    `gorm:"type:uuid" json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for Inventory
func (Inventory) TableName() string {
	return "inventories"
}
