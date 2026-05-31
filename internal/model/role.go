package model

import "time"

// Role represents a user role with permissions
type Role struct {
	ID          string       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string       `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Description string       `gorm:"type:varchar(255)" json:"description"`
	IsActive    bool         `gorm:"default:true" json:"is_active"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// TableName specifies the table name for Role
func (Role) TableName() string {
	return "roles"
}

// Permission represents a system permission
type Permission struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Module      string `gorm:"type:varchar(100);not null" json:"module"`
}

// TableName specifies the table name for Permission
func (Permission) TableName() string {
	return "permissions"
}
