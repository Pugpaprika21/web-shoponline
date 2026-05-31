package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user database model
type User struct {
	ID           string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex:idx_users_email;not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	FullName     string         `gorm:"type:varchar(255);not null" json:"full_name"`
	RoleID       *string        `gorm:"type:uuid" json:"role_id"`
	Role         *Role          `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}
