package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
)

func (s *DeviceService) Create(ctx context.Context, input dto.CreateDeviceInput) (*domain.Device, bool, error) {
	device := &domain.Device{
		Name:  input.Name,
		Brand: input.Brand,
		State: input.State,
	}

	return s.repo.Create(ctx, device)
}
