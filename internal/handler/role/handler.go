package role

import "web-shoponline/internal/service"

// Handler handles HTTP requests for role management
type Handler struct {
	roleService service.RoleServiceInterface
}

// NewHandler creates a new role Handler
func NewHandler(roleService service.RoleServiceInterface) *Handler {
	return &Handler{roleService: roleService}
}
