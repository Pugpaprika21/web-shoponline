package menu

import "context"

// DeleteMenu removes a menu
func (s *Service) DeleteMenu(ctx context.Context, id string) error {
	return s.menuRepo.Delete(ctx, id)
}
