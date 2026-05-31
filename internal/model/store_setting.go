package model

import "time"

// StoreSetting represents a key-value store setting
type StoreSetting struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Key       string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"key"`
	Value     string    `gorm:"type:text" json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for StoreSetting
func (StoreSetting) TableName() string {
	return "store_settings"
}
