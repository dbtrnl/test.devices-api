package usecase

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
)

func (s *DeviceService) DeleteByExternalID(ctx context.Context, input dto.DeviceUUIDInput) error {
	return s.repo.DeleteByExternalID(ctx, input.ExternalID)
}
