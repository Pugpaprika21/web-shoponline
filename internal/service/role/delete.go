package role

import "context"

// DeleteRole removes a role
func (s *Service) DeleteRole(ctx context.Context, id string) error {
	return s.roleRepo.Delete(ctx, id)
}
