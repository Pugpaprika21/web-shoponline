package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"web-shoponline/internal/model"
)

// Seed populates the database with initial data
func Seed(db *gorm.DB) error {
	log.Println("Seeding database...")

	// Seed permissions
	permissions := []model.Permission{
		{Name: "products.view", Description: "View products", Module: "products"},
		{Name: "products.create", Description: "Create products", Module: "products"},
		{Name: "products.edit", Description: "Edit products", Module: "products"},
		{Name: "products.delete", Description: "Delete products", Module: "products"},
		{Name: "orders.view", Description: "View orders", Module: "orders"},
		{Name: "orders.manage", Description: "Manage orders", Module: "orders"},
		{Name: "inventory.view", Description: "View inventory", Module: "inventory"},
		{Name: "inventory.manage", Description: "Manage inventory", Module: "inventory"},
		{Name: "users.view", Description: "View users", Module: "users"},
		{Name: "users.manage", Description: "Manage users", Module: "users"},
		{Name: "roles.manage", Description: "Manage roles", Module: "roles"},
		{Name: "menus.manage", Description: "Manage menus", Module: "menus"},
		{Name: "settings.manage", Description: "Manage settings", Module: "settings"},
	}

	for i := range permissions {
		db.Where("name = ?", permissions[i].Name).FirstOrCreate(&permissions[i])
	}

	// Seed roles
	adminRole := model.Role{Name: "admin", Description: "System Administrator"}
	db.Where("name = ?", adminRole.Name).FirstOrCreate(&adminRole)
	db.Model(&adminRole).Association("Permissions").Replace(permissions)

	customerRole := model.Role{Name: "customer", Description: "Customer"}
	db.Where("name = ?", customerRole.Name).FirstOrCreate(&customerRole)

	// Seed admin user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminUser := model.User{
		Email:        "admin@shoponline.com",
		PasswordHash: string(hashedPassword),
		FullName:     "Admin User",
		RoleID:       &adminRole.ID,
		IsActive:     true,
	}
	db.Where("email = ?", adminUser.Email).FirstOrCreate(&adminUser)

	// Seed categories
	categories := []model.Category{
		{Name: "Electronics", Slug: "electronics", Description: "Electronic devices and gadgets"},
		{Name: "Clothing", Slug: "clothing", Description: "Fashion and apparel"},
		{Name: "Books", Slug: "books", Description: "Books and publications"},
		{Name: "Home & Garden", Slug: "home-garden", Description: "Home and garden products"},
	}
	for i := range categories {
		db.Where("slug = ?", categories[i].Slug).FirstOrCreate(&categories[i])
	}

	// Seed menus
	menus := []model.Menu{
		{Name: "Dashboard", Icon: "dashboard", URL: "/admin", SortOrder: 1, IsActive: true},
		{Name: "Products", Icon: "inventory_2", URL: "/admin/products", SortOrder: 2, IsActive: true},
		{Name: "Inventory", Icon: "warehouse", URL: "/admin/inventory", SortOrder: 3, IsActive: true},
		{Name: "Orders", Icon: "shopping_cart", URL: "/admin/orders", SortOrder: 4, IsActive: true},
		{Name: "Users", Icon: "people", URL: "/admin/users", SortOrder: 5, IsActive: true},
		{Name: "Settings", Icon: "settings", URL: "#", SortOrder: 10, IsActive: true},
	}
	for i := range menus {
		db.Where("name = ? AND parent_id IS NULL", menus[i].Name).FirstOrCreate(&menus[i])
	}

	// Seed sub-menus under Settings
	var settingsMenu model.Menu
	db.Where("name = ? AND url = ?", "Settings", "#").First(&settingsMenu)
	if settingsMenu.ID != "" {
		subMenus := []model.Menu{
			{ParentID: &settingsMenu.ID, Name: "Menus", Icon: "menu", URL: "/admin/settings/menus", SortOrder: 1, IsActive: true},
			{ParentID: &settingsMenu.ID, Name: "Roles", Icon: "admin_panel_settings", URL: "/admin/settings/roles", SortOrder: 2, IsActive: true},
			{ParentID: &settingsMenu.ID, Name: "Store", Icon: "storefront", URL: "/admin/settings/store", SortOrder: 3, IsActive: true},
		}
		for i := range subMenus {
			db.Where("name = ? AND parent_id = ?", subMenus[i].Name, subMenus[i].ParentID).FirstOrCreate(&subMenus[i])
		}
	}

	// Seed default store settings
	defaultSettings := map[string]string{
		"store_name":       "ShopOnline",
		"store_logo_url":   "",
		"store_banner_url": "",
		"primary_color":    "#4F46E5",
		"secondary_color":  "#7C3AED",
		"footer_text":      "© 2024 ShopOnline. All rights reserved.",
		"social_facebook":  "",
		"social_line":      "",
		"social_instagram": "",
		"contact_phone":    "",
		"contact_email":    "",
		"contact_address":  "",
	}
	for key, value := range defaultSettings {
		setting := model.StoreSetting{Key: key, Value: value}
		db.Where("key = ?", key).FirstOrCreate(&setting)
	}

	log.Println("Database seeding completed")
	return nil
}
