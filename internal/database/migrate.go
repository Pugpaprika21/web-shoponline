package database

import (
	"log"

	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// AutoMigrate runs GORM auto-migration for all models
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&model.Permission{},
		&model.Role{},
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Inventory{},
		&model.Cart{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{},
		&model.Menu{},
		&model.StoreSetting{},
	)
	if err != nil {
		return err
	}

	log.Println("Database migrations completed")
	return nil
}
