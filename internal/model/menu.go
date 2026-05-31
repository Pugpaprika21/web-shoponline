package model

import "time"

// Menu represents a navigation menu item for admin panel
type Menu struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ParentID  *string   `gorm:"type:uuid" json:"parent_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Icon      string    `gorm:"type:varchar(100)" json:"icon"`
	URL       string    `gorm:"type:varchar(255)" json:"url"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	RoleIDs   []Role    `gorm:"many2many:menu_roles;" json:"roles,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Menu
func (Menu) TableName() string {
	return "menus"
}
