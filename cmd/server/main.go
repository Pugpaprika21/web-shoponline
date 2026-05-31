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
	"web-shoponline/internal/handler"
	"web-shoponline/internal/repository"
	"web-shoponline/internal/router"
	"web-shoponline/internal/service"
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
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	cartRepo := repository.NewCartRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	storeSettingRepo := repository.NewStoreSettingRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, roleRepo)
	productService := service.NewProductService(productRepo, categoryRepo)
	orderService := service.NewOrderService(orderRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	inventoryService := service.NewInventoryService(inventoryRepo, productRepo)
	menuService := service.NewMenuService(menuRepo)
	roleService := service.NewRoleService(roleRepo)
	userService := service.NewUserService(userRepo)
	storeSettingService := service.NewStoreSettingService(storeSettingRepo)

	// Initialize handlers
	handlers := &router.Handlers{
		Page:         handler.NewPageHandler(),
		Auth:         handler.NewAuthHandler(authService, cfg.JWT),
		Product:      handler.NewProductHandler(productService),
		Cart:         handler.NewCartHandler(cartService),
		Order:        handler.NewOrderHandler(orderService),
		Inventory:    handler.NewInventoryHandler(inventoryService, productService),
		Menu:         handler.NewMenuHandler(menuService),
		Role:         handler.NewRoleHandler(roleService),
		User:         handler.NewUserHandler(userService, roleService),
		StoreSetting: handler.NewStoreSettingHandler(storeSettingService),
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
