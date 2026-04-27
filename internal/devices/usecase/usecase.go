package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
)

type deviceRepository interface {
	GetByExternalID(ctx context.Context, id string) (*domain.Device, error)
	Create(ctx context.Context, device *domain.Device) (*domain.Device, error)
	List(ctx context.Context, filter domain.DeviceFilter) ([]domain.Device, error)
	DeleteByExternalID(ctx context.Context, id string) error
	Update(ctx context.Context, device domain.UpdateDevice) (*domain.Device, error)
}

type DeviceService struct {
	repo deviceRepository
}

func NewDeviceService(repo deviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}
