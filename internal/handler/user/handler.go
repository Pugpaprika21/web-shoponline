package user

import "web-shoponline/internal/service"

// Handler handles HTTP requests for user management
type Handler struct {
	userService service.UserServiceInterface
	roleService service.RoleServiceInterface
}

// NewHandler creates a new user Handler
func NewHandler(userService service.UserServiceInterface, roleService service.RoleServiceInterface) *Handler {
	return &Handler{
		userService: userService,
		roleService: roleService,
	}
}
