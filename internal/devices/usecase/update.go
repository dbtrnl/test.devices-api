package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
)

func (s *DeviceService) Update(ctx context.Context, req dto.UpdateDeviceInput) (*domain.Device, error) {
	var device domain.UpdateDevice

	device.ExternalID = req.ExternalID

	if req.Name != nil {
		device.Name = *req.Name
	}
	if req.Brand != nil {
		device.Brand = *req.Brand
	}
	if req.State != nil {
		device.State = domain.DeviceState(*req.State)
	}

	return s.repo.Update(ctx, device)
}
