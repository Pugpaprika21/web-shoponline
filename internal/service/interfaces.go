package service

import (
	"context"

	"web-shoponline/internal/dto"
)

// ProductServiceInterface defines the contract for product business logic
type ProductServiceInterface interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error)
	GetAllProductsPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error)
	GetAllProductsAdmin(ctx context.Context) ([]dto.ProductResponse, error)
	GetAllProductsAdminPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.ProductResponse, dto.PaginationResponse, error)
	GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error)
	CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, id string) error
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error)
}

// AuthServiceInterface defines the contract for authentication business logic
type AuthServiceInterface interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserResponse, error)
}

// OrderServiceInterface defines the contract for order business logic
type OrderServiceInterface interface {
	GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error)
	GetAllOrdersPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.OrderResponse, dto.PaginationResponse, error)
	GetOrderByID(ctx context.Context, id string) (*dto.OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, id string, req dto.UpdateOrderStatusRequest) error
}

// CartServiceInterface defines the contract for cart business logic
type CartServiceInterface interface {
	GetCart(ctx context.Context, userID string) (*dto.CartResponse, error)
	AddToCart(ctx context.Context, userID string, req dto.AddToCartRequest) error
	RemoveFromCart(ctx context.Context, userID, itemID string) error
	GetCartItemCount(ctx context.Context, userID string) (int64, error)
}

// InventoryServiceInterface defines the contract for inventory business logic
type InventoryServiceInterface interface {
	GetAll(ctx context.Context) ([]dto.InventoryResponse, error)
	GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.InventoryResponse, dto.PaginationResponse, error)
	AddInventory(ctx context.Context, req dto.CreateInventoryRequest, createdBy string) (*dto.InventoryResponse, error)
}

// MenuServiceInterface defines the contract for menu business logic
type MenuServiceInterface interface {
	GetAllMenus(ctx context.Context) ([]dto.MenuResponse, error)
	GetAllMenusPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.MenuResponse, dto.PaginationResponse, error)
	CreateMenu(ctx context.Context, req dto.CreateMenuRequest) (*dto.MenuResponse, error)
	UpdateMenu(ctx context.Context, id string, req dto.UpdateMenuRequest) error
	DeleteMenu(ctx context.Context, id string) error
}

// RoleServiceInterface defines the contract for role business logic
type RoleServiceInterface interface {
	GetAllRoles(ctx context.Context) ([]dto.RoleResponse, error)
	GetAllRolesPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.RoleResponse, dto.PaginationResponse, error)
	GetAllPermissions(ctx context.Context) ([]dto.PermissionResponse, error)
	CreateRole(ctx context.Context, req dto.CreateRoleRequest) (*dto.RoleResponse, error)
	UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, id string) error
}

// UserServiceInterface defines the contract for user management business logic
type UserServiceInterface interface {
	GetAllPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.AdminUserResponse, dto.PaginationResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.AdminUserResponse, error)
	UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) error
}

// StoreSettingServiceInterface defines the contract for store setting business logic
type StoreSettingServiceInterface interface {
	GetAll(ctx context.Context) (dto.StoreSettingsResponse, error)
	GetAllAsMap(ctx context.Context) (map[string]string, error)
	UpdateAll(ctx context.Context, req dto.StoreSettingsRequest) error
}
