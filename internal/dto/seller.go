package dto

// SellerCreateProductRequest is the DTO for creating a product by the seller
type SellerCreateProductRequest struct {
	CategoryID  string  `json:"category_id" form:"category_id" validate:"omitempty,uuid"`
	Name        string  `json:"name" form:"name" validate:"required,min=2,max=255"`
	Description string  `json:"description" form:"description" validate:"omitempty,max=2000"`
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock_quantity" form:"stock_quantity" validate:"required,gte=0"`
	ImageURL    string  `json:"image_url" form:"image_url" validate:"omitempty"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}

// SellerUpdateProductRequest is the DTO for updating a product by the seller
type SellerUpdateProductRequest struct {
	CategoryID  string  `json:"category_id" form:"category_id" validate:"omitempty,uuid"`
	Name        string  `json:"name" form:"name" validate:"omitempty,min=2,max=255"`
	Description string  `json:"description" form:"description" validate:"omitempty,max=2000"`
	Price       float64 `json:"price" form:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock_quantity" form:"stock_quantity" validate:"omitempty,gte=0"`
	ImageURL    string  `json:"image_url" form:"image_url" validate:"omitempty"`
	IsActive    *bool   `json:"is_active" form:"is_active"`
}

// UpdateShopRequest is the DTO for updating shop info by the seller
type UpdateShopRequest struct {
	Name        string `json:"name" form:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" form:"description" validate:"omitempty,max=2000"`
	LogoURL     string `json:"logo_url" form:"logo_url" validate:"omitempty"`
	BannerURL   string `json:"banner_url" form:"banner_url" validate:"omitempty"`
	Phone       string `json:"phone" form:"phone" validate:"omitempty,max=50"`
	Email       string `json:"email" form:"email" validate:"omitempty,email,max=255"`
	Address     string `json:"address" form:"address" validate:"omitempty,max=1000"`
}

// CreateShopRequest is the DTO for creating a new shop by the seller
type CreateShopRequest struct {
	Name        string `json:"name" form:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" form:"description" validate:"omitempty,max=2000"`
}
