package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
)

// func (s *DeviceService) List(ctx context.Context, filter domain.DeviceFilter) ([]domain.Device, error) {
// 	return s.repo.List(ctx, filter)
// }

func (uc *DeviceService) List(ctx context.Context, input dto.ListDevicesInput) ([]dto.DeviceResponse, error) {

	var filter domain.DeviceFilter

	if input.Brand != nil {
		filter.Brand = input.Brand
	}

	if input.State != nil {
		s := domain.DeviceState(*input.State)
		filter.State = &s
	}

	devices, err := uc.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	output := make([]dto.DeviceResponse, 0, len(devices))
	for _, d := range devices {
		output = append(output, dto.ToDeviceResponse(&d))
	}
	

	return output, nil
}
