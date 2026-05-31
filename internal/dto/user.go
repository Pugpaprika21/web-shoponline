package dto

// CreateUserRequest is the DTO for creating a user
type CreateUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	FullName string `json:"full_name" form:"full_name" validate:"required,min=2,max=255"`
	RoleID   string `json:"role_id" form:"role_id" validate:"required,uuid"`
}

// UpdateUserRequest is the DTO for updating a user
type UpdateUserRequest struct {
	FullName string `json:"full_name" form:"full_name" validate:"omitempty,min=2,max=255"`
	RoleID   string `json:"role_id" form:"role_id" validate:"omitempty,uuid"`
	IsActive *bool  `json:"is_active" form:"is_active"`
}

// AdminUserResponse is the DTO for returning user data in admin context
type AdminUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	RoleID    string `json:"role_id"`
	RoleName  string `json:"role_name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}
