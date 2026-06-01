package model

import (
	"time"

	"gorm.io/gorm"
)

// Shop represents a store/shop in the system
type Shop struct {
	ID          string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Slug        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description string         `gorm:"type:text" json:"description"`
	LogoURL     string         `gorm:"type:varchar(500)" json:"logo_url"`
	BannerURL   string         `gorm:"type:varchar(500)" json:"banner_url"`
	OwnerID     string         `gorm:"type:uuid;not null" json:"owner_id"`
	Owner       User           `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Phone       string         `gorm:"type:varchar(50)" json:"phone"`
	Email       string         `gorm:"type:varchar(255)" json:"email"`
	Address     string         `gorm:"type:text" json:"address"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	IsVerified  bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Shop
func (Shop) TableName() string {
	return "shops"
}
