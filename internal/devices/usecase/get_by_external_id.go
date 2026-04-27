package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
)

func (s *DeviceService) GetByExternalID(ctx context.Context, input dto.DeviceUUIDInput) (*domain.Device, error) {
    return s.repo.GetByExternalID(ctx, input.ExternalID)
}
