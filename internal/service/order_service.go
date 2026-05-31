package service

import (
	"context"
	"fmt"
	"math"

	"web-shoponline/internal/dto"
	"web-shoponline/internal/repository"
)

// OrderService handles business logic for orders
type OrderService struct {
	orderRepo *repository.OrderRepository
}

// NewOrderService creates a new OrderService
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

// GetAllOrders retrieves all orders as DTOs
func (s *OrderService) GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	orders, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get orders: %w", err)
	}

	responses := make([]dto.OrderResponse, 0, len(orders))
	for _, o := range orders {
		resp := dto.OrderResponse{
			ID:              o.ID,
			Status:          o.Status,
			TotalAmount:     o.TotalAmount,
			ShippingAddress: o.ShippingAddress,
			CreatedAt:       o.CreatedAt,
		}
		if o.User.Email != "" {
			resp.UserEmail = o.User.Email
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// GetAllOrdersPaginated retrieves orders with pagination and optional search
func (s *OrderService) GetAllOrdersPaginated(ctx context.Context, page dto.PaginationRequest, search string) ([]dto.OrderResponse, dto.PaginationResponse, error) {
	page.Normalize()
	orders, total, err := s.orderRepo.GetAllPaginated(ctx, page.Offset(), page.PageSize, search)
	if err != nil {
		return nil, dto.PaginationResponse{}, fmt.Errorf("service: failed to get orders: %w", err)
	}

	responses := make([]dto.OrderResponse, 0, len(orders))
	for _, o := range orders {
		resp := dto.OrderResponse{
			ID:              o.ID,
			Status:          o.Status,
			TotalAmount:     o.TotalAmount,
			ShippingAddress: o.ShippingAddress,
			CreatedAt:       o.CreatedAt,
		}
		if o.User.Email != "" {
			resp.UserEmail = o.User.Email
		}
		responses = append(responses, resp)
	}

	pagination := dto.PaginationResponse{
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(page.PageSize))),
	}
	return responses, pagination, nil
}

// GetOrderByID retrieves a single order with items
func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*dto.OrderResponse, error) {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get order: %w", err)
	}

	resp := &dto.OrderResponse{
		ID:              order.ID,
		UserEmail:       order.User.Email,
		Status:          order.Status,
		TotalAmount:     order.TotalAmount,
		ShippingAddress: order.ShippingAddress,
		CreatedAt:       order.CreatedAt,
		Items:           make([]dto.OrderItemResponse, 0),
	}

	for _, item := range order.Items {
		resp.Items = append(resp.Items, dto.OrderItemResponse{
			ID:          item.ID,
			ProductName: item.Product.Name,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			Subtotal:    float64(item.Quantity) * item.UnitPrice,
		})
	}

	return resp, nil
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(ctx context.Context, id string, req dto.UpdateOrderStatusRequest) error {
	if err := s.orderRepo.UpdateStatus(ctx, id, req.Status); err != nil {
		return fmt.Errorf("service: failed to update order status: %w", err)
	}
	return nil
}
