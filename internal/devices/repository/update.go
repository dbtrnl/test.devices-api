package repository

import (
	"context"
	"errors"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/infra/dberrors"
	"gorm.io/gorm"
)

func (r *deviceRepository) Update(ctx context.Context, input domain.UpdateDevice) (*domain.Device, error) {
	var existing domain.Device

	err := r.db.WithContext(ctx).
		Where("external_id = ?", input.ExternalID).
		First(&existing).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.NewErrDeviceNotFound(input.ExternalID)
	}
	if err != nil {
		return nil, err
	}
	if existing.IsDeleted {
		return nil, domain.NewErrDeviceDeleted(input.ExternalID)
	}

	upd, err := updateDeviceValues(existing, input.Name, input.Brand, input.State)
	if err != nil {
		return nil, err
	}

	result := r.db.WithContext(ctx).Save(&upd)
	if result.Error != nil {
		return nil, dberrors.Translate(result.Error, input.ExternalID)
	}

	return &upd, nil
}

func updateDeviceValues(existingDevice domain.Device, name, brand string, state domain.DeviceState) (domain.Device, error) {
	newStateIsValid := state != domain.DeviceStateInUse
	deviceIsInUse := existingDevice.State == domain.DeviceStateInUse

	if deviceIsInUse && !newStateIsValid {
		return domain.Device{}, domain.NewErrDeviceInUse(existingDevice.ExternalID)
	}
	if name != "" {
		existingDevice.Name = name
	}
	if brand != "" {
		existingDevice.Brand = brand
	}
	if state != "" {
		existingDevice.State = state
	}
	return existingDevice, nil
}
