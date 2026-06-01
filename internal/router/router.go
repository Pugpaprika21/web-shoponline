package router

import (
	"github.com/gofiber/fiber/v2"

	"web-shoponline/internal/config"
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
	"web-shoponline/internal/middleware"
)

// Handlers holds all handler instances
type Handlers struct {
	Page         *handlerPage.Handler
	Auth         *handlerAuth.Handler
	Product      *handlerProduct.Handler
	Cart         *handlerCart.Handler
	Order        *handlerOrder.Handler
	Inventory    *handlerInventory.Handler
	Menu         *handlerMenu.Handler
	Role         *handlerRole.Handler
	User         *handlerUser.Handler
	StoreSetting *handlerStoreSetting.Handler
	Shop         *handlerShop.Handler
	Seller       *handlerSeller.Handler
}

// Setup configures all application routes
func Setup(app *fiber.App, h *Handlers, jwtCfg config.JWTConfig) {
	// Static files
	app.Static("/static", "./web/static")

	// --- Public Pages ---
	app.Get("/", h.Page.RenderHomePage)
	app.Get("/products", h.Product.RenderProductsPage)
	app.Get("/products/:id", h.Product.RenderProductDetailPage)
	app.Get("/shops/:slug", h.Shop.RenderShopPage)
	app.Get("/login", h.Auth.RenderLoginPage)
	app.Get("/register", h.Auth.RenderRegisterPage)
	app.Get("/cart", h.Cart.RenderCartPage)

	// --- Admin Login (no auth required) ---
	app.Get("/admin/login", h.Auth.RenderAdminLoginPage)

	// --- Public API (no auth) ---
	api := app.Group("/api")
	api.Get("/products", h.Product.GetAllProducts)
	api.Get("/products/:id", h.Product.GetProductByID)
	api.Get("/store-info", h.StoreSetting.GetPublicStoreInfo)

	// Auth API (no auth required)
	auth := api.Group("/auth")
	auth.Post("/login", h.Auth.Login)
	auth.Post("/register", h.Auth.Register)
	auth.Get("/me", h.Auth.GetCurrentUser)

	// Auth API (requires JWT)
	authProtected := api.Group("/auth", middleware.RequireJWT(jwtCfg))
	authProtected.Post("/logout", h.Auth.Logout)

	// Cart API (requires JWT auth)
	cart := api.Group("/cart", middleware.RequireJWT(jwtCfg))
	cart.Get("/", h.Cart.GetCart)
	cart.Post("/items", h.Cart.AddToCart)
	cart.Delete("/items/:id", h.Cart.RemoveFromCart)
	cart.Get("/count", h.Cart.GetCartCount)

	// --- Admin Pages (requires admin JWT) ---
	admin := app.Group("/admin", middleware.RequireAdminJWT(jwtCfg))
	admin.Get("/", h.Page.RenderAdminDashboard)
	admin.Get("/products", h.Product.RenderAdminProductsPage)
	admin.Get("/inventory", h.Inventory.RenderAdminInventoryPage)
	admin.Get("/orders", h.Order.RenderAdminOrdersPage)
	admin.Get("/settings/menus", h.Menu.RenderAdminMenusPage)
	admin.Get("/settings/roles", h.Role.RenderAdminRolesPage)
	admin.Get("/settings/store", h.StoreSetting.RenderAdminStoreSettingsPage)
	admin.Get("/users", h.User.RenderAdminUsersPage)

	// --- Admin API (requires admin JWT) ---
	adminAPI := api.Group("/admin", middleware.RequireAdminJWT(jwtCfg))

	// Products
	adminAPI.Get("/products", h.Product.GetAllProducts)
	adminAPI.Post("/products", h.Product.CreateProduct)
	adminAPI.Put("/products/:id", h.Product.UpdateProduct)
	adminAPI.Delete("/products/:id", h.Product.DeleteProduct)

	// Inventory
	adminAPI.Get("/inventory", h.Inventory.GetAll)
	adminAPI.Post("/inventory", h.Inventory.AddInventory)

	// Orders
	adminAPI.Get("/orders", h.Order.GetAllOrders)
	adminAPI.Get("/orders/:id", h.Order.GetOrderByID)
	adminAPI.Patch("/orders/:id/status", h.Order.UpdateOrderStatus)

	// Menus
	adminAPI.Get("/menus", h.Menu.GetAllMenus)
	adminAPI.Post("/menus", h.Menu.CreateMenu)
	adminAPI.Put("/menus/:id", h.Menu.UpdateMenu)
	adminAPI.Delete("/menus/:id", h.Menu.DeleteMenu)

	// Roles
	adminAPI.Get("/roles", h.Role.GetAllRoles)
	adminAPI.Get("/permissions", h.Role.GetAllPermissions)
	adminAPI.Post("/roles", h.Role.CreateRole)
	adminAPI.Put("/roles/:id", h.Role.UpdateRole)
	adminAPI.Delete("/roles/:id", h.Role.DeleteRole)

	// Users
	adminAPI.Get("/users", h.User.GetAllUsers)
	adminAPI.Post("/users", h.User.CreateUser)
	adminAPI.Put("/users/:id", h.User.UpdateUser)

	// Store Settings
	adminAPI.Get("/settings/store", h.StoreSetting.GetStoreSettings)
	adminAPI.Put("/settings/store", h.StoreSetting.UpdateStoreSettings)

	// File Upload
	adminAPI.Post("/upload", h.StoreSetting.UploadImage)

	// --- Seller Pages (requires seller JWT) ---
	seller := app.Group("/seller", middleware.RequireSellerJWT(jwtCfg))
	seller.Get("/", h.Seller.RenderSellerDashboard)
	seller.Get("/shop", h.Seller.RenderSellerShopSettings)
	seller.Get("/products", h.Seller.RenderSellerProducts)
	seller.Get("/orders", h.Seller.RenderSellerOrders)

	// --- Seller API (requires seller JWT) ---
	sellerAPI := api.Group("/seller", middleware.RequireSellerJWT(jwtCfg))
	sellerAPI.Get("/shop", h.Seller.GetMyShop)
	sellerAPI.Post("/shop", h.Seller.CreateMyShop)
	sellerAPI.Put("/shop", h.Seller.UpdateMyShop)
	sellerAPI.Get("/products", h.Seller.GetMyProducts)
	sellerAPI.Post("/products", h.Seller.CreateProduct)
	sellerAPI.Put("/products/:id", h.Seller.UpdateProduct)
	sellerAPI.Delete("/products/:id", h.Seller.DeleteProduct)
	sellerAPI.Get("/orders", h.Seller.GetMyOrders)

	// Seller file upload (reuse same upload logic)
	sellerAPI.Post("/upload", h.StoreSetting.UploadImage)
}
