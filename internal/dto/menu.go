package dto

// CreateMenuRequest is the DTO for creating a menu
type CreateMenuRequest struct {
	ParentID  string `json:"parent_id" form:"parent_id" validate:"omitempty,uuid"`
	Name      string `json:"name" form:"name" validate:"required,min=2,max=100"`
	Icon      string `json:"icon" form:"icon" validate:"omitempty,max=100"`
	URL       string `json:"url" form:"url" validate:"required,max=255"`
	SortOrder int    `json:"sort_order" form:"sort_order" validate:"gte=0"`
	IsActive  bool   `json:"is_active" form:"is_active"`
}

// UpdateMenuRequest is the DTO for updating a menu
type UpdateMenuRequest struct {
	ParentID  string `json:"parent_id" form:"parent_id" validate:"omitempty,uuid"`
	Name      string `json:"name" form:"name" validate:"omitempty,min=2,max=100"`
	Icon      string `json:"icon" form:"icon" validate:"omitempty,max=100"`
	URL       string `json:"url" form:"url" validate:"omitempty,max=255"`
	SortOrder int    `json:"sort_order" form:"sort_order" validate:"gte=0"`
	IsActive  *bool  `json:"is_active" form:"is_active"`
}

// MenuResponse is the DTO for returning menu data
type MenuResponse struct {
	ID        string         `json:"id"`
	ParentID  string         `json:"parent_id,omitempty"`
	Name      string         `json:"name"`
	Icon      string         `json:"icon"`
	URL       string         `json:"url"`
	SortOrder int            `json:"sort_order"`
	IsActive  bool           `json:"is_active"`
	Children  []MenuResponse `json:"children,omitempty"`
}
