package dto

// ProductResponse is the DTO for returning product data to the client
type ProductResponse struct {
	ID            string  `json:"id"`
	CategoryID    string  `json:"category_id,omitempty"`
	CategoryName  string  `json:"category_name,omitempty"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	ImageURL      string  `json:"image_url,omitempty"`
	IsActive      bool    `json:"is_active"`
}

// CreateProductRequest is the DTO for creating a new product
type CreateProductRequest struct {
	CategoryID  string  `json:"category_id" form:"category_id" validate:"omitempty,uuid"`
	Name        string  `json:"name" form:"name" validate:"required,min=2,max=255"`
	Description string  `json:"description" form:"description" validate:"omitempty,max=2000"`
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock_quantity" form:"stock_quantity" validate:"required,gte=0"`
	ImageURL    string  `json:"image_url" form:"image_url" validate:"omitempty,url"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}

// UpdateProductRequest is the DTO for updating an existing product
type UpdateProductRequest struct {
	CategoryID  string  `json:"category_id" form:"category_id" validate:"omitempty,uuid"`
	Name        string  `json:"name" form:"name" validate:"omitempty,min=2,max=255"`
	Description string  `json:"description" form:"description" validate:"omitempty,max=2000"`
	Price       float64 `json:"price" form:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock_quantity" form:"stock_quantity" validate:"omitempty,gte=0"`
	ImageURL    string  `json:"image_url" form:"image_url" validate:"omitempty,url"`
	IsActive    *bool   `json:"is_active" form:"is_active"`
}

// CategoryResponse is the DTO for returning category data
type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

// CreateCategoryRequest is the DTO for creating a new category
type CreateCategoryRequest struct {
	Name        string `json:"name" form:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" form:"description" validate:"omitempty,max=500"`
}
