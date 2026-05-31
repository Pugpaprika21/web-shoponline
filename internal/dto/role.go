package dto

// CreateRoleRequest is the DTO for creating a role
type CreateRoleRequest struct {
	Name          string   `json:"name" form:"name" validate:"required,min=2,max=100"`
	Description   string   `json:"description" form:"description" validate:"omitempty,max=255"`
	PermissionIDs []string `json:"permission_ids" form:"permission_ids" validate:"omitempty"`
}

// UpdateRoleRequest is the DTO for updating a role
type UpdateRoleRequest struct {
	Name          string   `json:"name" form:"name" validate:"omitempty,min=2,max=100"`
	Description   string   `json:"description" form:"description" validate:"omitempty,max=255"`
	PermissionIDs []string `json:"permission_ids" form:"permission_ids" validate:"omitempty"`
	IsActive      *bool    `json:"is_active" form:"is_active"`
}

// RoleResponse is the DTO for returning role data
type RoleResponse struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	IsActive    bool                 `json:"is_active"`
	Permissions []PermissionResponse `json:"permissions,omitempty"`
}

// PermissionResponse is the DTO for returning permission data
type PermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Module      string `json:"module"`
}
