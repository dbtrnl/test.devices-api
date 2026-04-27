package repository

import (
	"context"
	"errors"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"gorm.io/gorm"
)

func (r *deviceRepository) GetByExternalID(ctx context.Context, id string) (*domain.Device, error) {
	var device domain.Device

	err := r.db.WithContext(ctx).
		First(&device, "external_id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.NewErrDeviceNotFound(id)
	}
	if err != nil {
		return nil, err
	}

	return &device, nil
}