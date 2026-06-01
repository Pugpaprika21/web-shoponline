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

	sellerRole := model.Role{Name: "seller", Description: "Shop Seller"}
	db.Where("name = ?", sellerRole.Name).FirstOrCreate(&sellerRole)

	// Seed admin user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminUser := model.User{
		Email:        "admin@shoponline.com",
		PasswordHash: string(hashedPassword),
		FullName:     "Admin User",
		IsActive:     true,
	}
	db.Where("email = ?", adminUser.Email).FirstOrCreate(&adminUser)
	db.Model(&adminUser).Association("Roles").Replace([]model.Role{adminRole})

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

	// Seed default shop
	defaultShop := model.Shop{
		Name:        "ShopOnline Official",
		Slug:        "shoponline-official",
		Description: "Official store",
		OwnerID:     adminUser.ID,
		IsActive:    true,
		IsVerified:  true,
	}
	db.Where("slug = ?", defaultShop.Slug).FirstOrCreate(&defaultShop)

	// Seed sample data for dashboard graphs
	seedSampleData(db, categories, adminUser, defaultShop)

	return nil
}

// seedSampleData creates sample products and orders for dashboard visualization
func seedSampleData(db *gorm.DB, categories []model.Category, adminUser model.User, defaultShop model.Shop) {
	// Check if sample data already exists
	var productCount int64
	db.Model(&model.Product{}).Count(&productCount)
	if productCount > 0 {
		return // Already seeded
	}

	log.Println("Seeding sample products and orders...")

	// Sample products
	products := []model.Product{
		{ShopID: defaultShop.ID, CategoryID: &categories[0].ID, Name: "iPhone 15 Pro", Slug: "iphone-15-pro", Description: "Latest Apple smartphone with A17 Pro chip", Price: 42900, StockQuantity: 50, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[0].ID, Name: "MacBook Air M3", Slug: "macbook-air-m3", Description: "Lightweight laptop with M3 chip", Price: 44900, StockQuantity: 30, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[0].ID, Name: "AirPods Pro 2", Slug: "airpods-pro-2", Description: "Active noise cancellation earbuds", Price: 8990, StockQuantity: 100, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[0].ID, Name: "iPad Air", Slug: "ipad-air", Description: "Versatile tablet for work and play", Price: 22900, StockQuantity: 40, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[0].ID, Name: "Samsung Galaxy S24", Slug: "samsung-galaxy-s24", Description: "AI-powered Android flagship", Price: 32900, StockQuantity: 45, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[1].ID, Name: "Nike Air Max 90", Slug: "nike-air-max-90", Description: "Classic running shoes", Price: 4500, StockQuantity: 80, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[1].ID, Name: "Levi's 501 Jeans", Slug: "levis-501-jeans", Description: "Original fit denim jeans", Price: 2990, StockQuantity: 60, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[1].ID, Name: "Uniqlo T-Shirt", Slug: "uniqlo-tshirt", Description: "Comfortable cotton t-shirt", Price: 390, StockQuantity: 200, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[1].ID, Name: "Adidas Hoodie", Slug: "adidas-hoodie", Description: "Warm fleece hoodie", Price: 2490, StockQuantity: 70, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[2].ID, Name: "Atomic Habits", Slug: "atomic-habits", Description: "Build good habits, break bad ones", Price: 450, StockQuantity: 150, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[2].ID, Name: "The Psychology of Money", Slug: "psychology-of-money", Description: "Timeless lessons on wealth", Price: 390, StockQuantity: 120, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[2].ID, Name: "Clean Code", Slug: "clean-code", Description: "A handbook of agile software craftsmanship", Price: 890, StockQuantity: 90, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[3].ID, Name: "Dyson V15 Vacuum", Slug: "dyson-v15-vacuum", Description: "Cordless vacuum cleaner", Price: 23900, StockQuantity: 25, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[3].ID, Name: "IKEA Desk Lamp", Slug: "ikea-desk-lamp", Description: "LED work lamp", Price: 790, StockQuantity: 110, IsActive: true},
		{ShopID: defaultShop.ID, CategoryID: &categories[3].ID, Name: "Plant Pot Set", Slug: "plant-pot-set", Description: "Ceramic pot set of 3", Price: 590, StockQuantity: 85, IsActive: true},
	}

	for i := range products {
		db.Create(&products[i])
	}

	// Sample customer users
	hashedPw, _ := bcrypt.GenerateFromPassword([]byte("customer123"), bcrypt.DefaultCost)
	var customerRole model.Role
	db.Where("name = ?", "customer").First(&customerRole)

	customers := []model.User{
		{Email: "somchai@example.com", PasswordHash: string(hashedPw), FullName: "Somchai Jaidee", Roles: []model.Role{customerRole}, IsActive: true},
		{Email: "somporn@example.com", PasswordHash: string(hashedPw), FullName: "Somporn Suksai", Roles: []model.Role{customerRole}, IsActive: true},
		{Email: "nattapong@example.com", PasswordHash: string(hashedPw), FullName: "Nattapong Kaewkla", Roles: []model.Role{customerRole}, IsActive: true},
		{Email: "pranee@example.com", PasswordHash: string(hashedPw), FullName: "Pranee Thongdee", Roles: []model.Role{customerRole}, IsActive: true},
		{Email: "wichai@example.com", PasswordHash: string(hashedPw), FullName: "Wichai Srisuwan", Roles: []model.Role{customerRole}, IsActive: true},
	}

	for i := range customers {
		db.Create(&customers[i])
	}

	// Sample orders with different statuses and dates
	statuses := []string{"pending", "confirmed", "shipped", "delivered", "delivered", "delivered", "cancelled"}
	addresses := []string{
		"123 Sukhumvit Rd, Bangkok 10110",
		"456 Ratchadaphisek Rd, Bangkok 10400",
		"789 Silom Rd, Bangkok 10500",
		"321 Phahonyothin Rd, Bangkok 10900",
		"654 Rama IV Rd, Bangkok 10330",
	}

	for i := 0; i < 20; i++ {
		customer := customers[i%len(customers)]
		status := statuses[i%len(statuses)]
		address := addresses[i%len(addresses)]

		// Pick 1-3 random products
		numItems := (i % 3) + 1
		var totalAmount float64
		var items []model.OrderItem

		for j := 0; j < numItems; j++ {
			product := products[(i+j)%len(products)]
			qty := (j % 3) + 1
			items = append(items, model.OrderItem{
				ProductID: product.ID,
				Quantity:  qty,
				UnitPrice: product.Price,
			})
			totalAmount += product.Price * float64(qty)
		}

		order := model.Order{
			UserID:          customer.ID,
			Status:          status,
			TotalAmount:     totalAmount,
			ShippingAddress: address,
			Items:           items,
		}
		db.Create(&order)
	}

	log.Println("Sample data seeded: 15 products, 5 customers, 20 orders")
}
