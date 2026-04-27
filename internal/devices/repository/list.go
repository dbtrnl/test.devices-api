package repository

import (
	"context"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
)

func (r *deviceRepository) List(ctx context.Context, filter domain.DeviceFilter) ([]domain.Device, error) {
	var devices []domain.Device
	db := r.db.WithContext(ctx).Model(&domain.Device{})

	if filter.Brand != nil {
		db = db.Where("brand = ?", *filter.Brand)
	}
	if filter.State != nil {
		db = db.Where("state = ?", *filter.State)
	}
	if err := db.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}