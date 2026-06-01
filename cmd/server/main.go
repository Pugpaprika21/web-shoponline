package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"web-shoponline/internal/config"
	"web-shoponline/internal/database"
	handlerAuth "web-shoponline/internal/handler/auth"
	handlerCart "web-shoponline/internal/handler/cart"
	handlerInventory "web-shoponline/internal/handler/inventory"
	handlerMenu "web-shoponline/internal/handler/menu"
	handlerOrder "web-shoponline/internal/handler/order"
	handlerPage "web-shoponline/internal/handler/page"
	handlerProduct "web-shoponline/internal/handler/product"
	handlerRole "web-shoponline/internal/handler/role"
	handlerSeller "web-shoponline/internal/handler/seller"
	handlerShop "web-shoponline/internal/handler/shop"
	handlerStoreSetting "web-shoponline/internal/handler/store_setting"
	handlerUser "web-shoponline/internal/handler/user"
	repoCart "web-shoponline/internal/repository/cart"
	repoCategory "web-shoponline/internal/repository/category"
	repoInventory "web-shoponline/internal/repository/inventory"
	repoMenu "web-shoponline/internal/repository/menu"
	repoOrder "web-shoponline/internal/repository/order"
	repoProduct "web-shoponline/internal/repository/product"
	repoRole "web-shoponline/internal/repository/role"
	repoShop "web-shoponline/internal/repository/shop"
	repoStoreSetting "web-shoponline/internal/repository/store_setting"
	repoUser "web-shoponline/internal/repository/user"
	"web-shoponline/internal/router"
	svcAuth "web-shoponline/internal/service/auth"
	svcCart "web-shoponline/internal/service/cart"
	svcInventory "web-shoponline/internal/service/inventory"
	svcMenu "web-shoponline/internal/service/menu"
	svcOrder "web-shoponline/internal/service/order"
	svcProduct "web-shoponline/internal/service/product"
	svcRole "web-shoponline/internal/service/role"
	svcStoreSetting "web-shoponline/internal/service/store_setting"
	svcUser "web-shoponline/internal/service/user"
)

func main() {
	// Load .env file (ignore error in production)
	_ = godotenv.Load()

	// Load configuration
	cfg := config.Load()

	// Connect to database using GORM
	db, err := database.NewGormDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Seed initial data
	if err := database.Seed(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	// Initialize template engine
	engine := html.New("./web/templates", ".html")
	if cfg.App.Env == "development" {
		engine.Reload(true)
	}
	engine.AddFunc("add", func(a, b int) int { return a + b })
	engine.AddFunc("subtract", func(a, b int) int { return a - b })
	engine.AddFunc("multiply", func(a, b int) int { return a * b })

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	// Initialize repositories
	userRepo := repoUser.NewRepository(db)
	productRepo := repoProduct.NewRepository(db)
	categoryRepo := repoCategory.NewRepository(db)
	orderRepo := repoOrder.NewRepository(db)
	cartRepo := repoCart.NewRepository(db)
	inventoryRepo := repoInventory.NewRepository(db)
	roleRepo := repoRole.NewRepository(db)
	menuRepo := repoMenu.NewRepository(db)
	storeSettingRepo := repoStoreSetting.NewRepository(db)
	shopRepo := repoShop.NewRepository(db)

	// Initialize services
	authService := svcAuth.NewService(userRepo, roleRepo)
	productService := svcProduct.NewService(productRepo, categoryRepo)
	orderService := svcOrder.NewService(orderRepo)
	cartService := svcCart.NewService(cartRepo, productRepo)
	inventoryService := svcInventory.NewService(inventoryRepo, productRepo)
	menuService := svcMenu.NewService(menuRepo)
	roleService := svcRole.NewService(roleRepo)
	userService := svcUser.NewService(userRepo, roleRepo)
	storeSettingService := svcStoreSetting.NewService(storeSettingRepo)

	// Initialize handlers
	handlers := &router.Handlers{
		Page:         handlerPage.NewHandler(),
		Auth:         handlerAuth.NewHandler(authService, cfg.JWT),
		Product:      handlerProduct.NewHandler(productService),
		Cart:         handlerCart.NewHandler(cartService),
		Order:        handlerOrder.NewHandler(orderService),
		Inventory:    handlerInventory.NewHandler(inventoryService, productService),
		Menu:         handlerMenu.NewHandler(menuService),
		Role:         handlerRole.NewHandler(roleService),
		User:         handlerUser.NewHandler(userService, roleService),
		StoreSetting: handlerStoreSetting.NewHandler(storeSettingService),
		Shop:         handlerShop.NewHandler(shopRepo, productRepo),
		Seller:       handlerSeller.NewHandler(shopRepo, productRepo, orderRepo),
	}

	// Setup routes
	router.Setup(app, handlers, cfg.JWT)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + cfg.App.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s (env: %s)", cfg.App.Port, cfg.App.Env)

	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}
