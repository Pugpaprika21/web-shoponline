package dto

// LoginRequest is the DTO for user login
type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

// RegisterRequest is the DTO for user registration
type RegisterRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	FullName string `json:"full_name" form:"full_name" validate:"required,min=2,max=255"`
}

// UserResponse is the DTO for returning user data
type UserResponse struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	FullName  string   `json:"full_name"`
	RoleNames []string `json:"role_names,omitempty"`
	IsActive  bool     `json:"is_active"`
}
