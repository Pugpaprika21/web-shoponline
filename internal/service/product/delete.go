package product

import "context"

// DeleteProduct removes a product
func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.Delete(ctx, id)
}
