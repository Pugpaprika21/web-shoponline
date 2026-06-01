package order

import (
	"context"
	"fmt"

	"web-shoponline/internal/dto"
)

// UpdateOrderStatus updates the status of an order
func (s *Service) UpdateOrderStatus(ctx context.Context, id string, req dto.UpdateOrderStatusRequest) error {
	if err := s.orderRepo.UpdateStatus(ctx, id, req.Status); err != nil {
		return fmt.Errorf("service: failed to update order status: %w", err)
	}
	return nil
}
