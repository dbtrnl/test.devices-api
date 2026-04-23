package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
)

func (s *DeviceService) GetByExternalID(ctx context.Context, id string) (*domain.Device, error) {
    return s.repo.GetByExternalID(ctx, id)
}