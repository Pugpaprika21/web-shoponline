package repository

import (
	"context"

	"web-shoponline/internal/model"
)

// ProductRepositoryInterface defines the contract for product data access
type ProductRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error)
	GetAllAdmin(ctx context.Context) ([]model.Product, error)
	GetAllAdminPaginated(ctx context.Context, offset, limit int, search string) ([]model.Product, int64, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	GetByShopID(ctx context.Context, shopID string) ([]model.Product, error)
	Create(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
}

// ShopRepositoryInterface defines the contract for shop data access
type ShopRepositoryInterface interface {
	GetByID(ctx context.Context, id string) (*model.Shop, error)
	GetBySlug(ctx context.Context, slug string) (*model.Shop, error)
	GetByOwnerID(ctx context.Context, ownerID string) (*model.Shop, error)
	GetAll(ctx context.Context) ([]model.Shop, error)
	Create(ctx context.Context, shop *model.Shop) error
	Update(ctx context.Context, shop *model.Shop) error
}

// CategoryRepositoryInterface defines the contract for category data access
type CategoryRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByID(ctx context.Context, id string) (*model.Category, error)
	Create(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
}

// OrderRepositoryInterface defines the contract for order data access
type OrderRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Order, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Order, int64, error)
	GetByID(ctx context.Context, id string) (*model.Order, error)
	GetByUserID(ctx context.Context, userID string) ([]model.Order, error)
	GetByShopID(ctx context.Context, shopID string) ([]model.Order, error)
	Create(ctx context.Context, order *model.Order) error
	UpdateStatus(ctx context.Context, id, status string) error
}

// UserRepositoryInterface defines the contract for user data access
type UserRepositoryInterface interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.User, int64, error)
	Update(ctx context.Context, user *model.User) error
	ReplaceRoles(ctx context.Context, user *model.User, roles []model.Role) error
}

// CartRepositoryInterface defines the contract for cart data access
type CartRepositoryInterface interface {
	GetOrCreateByUserID(ctx context.Context, userID string) (*model.Cart, error)
	AddItem(ctx context.Context, cartID, productID string, quantity int) error
	RemoveItem(ctx context.Context, cartID, itemID string) error
	ClearCart(ctx context.Context, cartID string) error
	GetItemCount(ctx context.Context, userID string) (int64, error)
}

// InventoryRepositoryInterface defines the contract for inventory data access
type InventoryRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Inventory, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Inventory, int64, error)
	GetByProductID(ctx context.Context, productID string) ([]model.Inventory, error)
	Create(ctx context.Context, record *model.Inventory) error
}

// RoleRepositoryInterface defines the contract for role data access
type RoleRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Role, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Role, int64, error)
	GetByID(ctx context.Context, id string) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	UpdatePermissions(ctx context.Context, roleID string, permissions []model.Permission) error
	Delete(ctx context.Context, id string) error
	GetAllPermissions(ctx context.Context) ([]model.Permission, error)
}

// MenuRepositoryInterface defines the contract for menu data access
type MenuRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Menu, error)
	GetAllPaginated(ctx context.Context, offset, limit int, search string) ([]model.Menu, int64, error)
	GetTopLevel(ctx context.Context) ([]model.Menu, error)
	GetChildren(ctx context.Context, parentID string) ([]model.Menu, error)
	GetByID(ctx context.Context, id string) (*model.Menu, error)
	Create(ctx context.Context, menu *model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, id string) error
}

// StoreSettingRepositoryInterface defines the contract for store setting data access
type StoreSettingRepositoryInterface interface {
	GetAll(ctx context.Context) (map[string]string, error)
	GetByKey(ctx context.Context, key string) (string, error)
	UpsertAll(ctx context.Context, settings map[string]string) error
}
