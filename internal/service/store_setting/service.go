package storesetting

import "web-shoponline/internal/repository"

// Service handles business logic for store settings
type Service struct {
	settingRepo repository.StoreSettingRepositoryInterface
}

// NewService creates a new store setting Service
func NewService(settingRepo repository.StoreSettingRepositoryInterface) *Service {
	return &Service{settingRepo: settingRepo}
}
