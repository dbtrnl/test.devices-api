package repository

import (
	"context"
	"errors"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *deviceRepository) Create(ctx context.Context, device *domain.Device) (*domain.Device, bool, error) {
	var d domain.Device
	var hasDeviceBeenCreated bool

	err := r.db.WithContext(ctx).
		Where("name = ? AND brand = ?", device.Name, device.Brand).
		First(&d).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, hasDeviceBeenCreated, err
	}
	if err == nil {
		if d.IsDeleted {
			return nil, hasDeviceBeenCreated, domain.NewErrDeviceExistsDeleted(d.Name, d.Brand, d.ExternalID)
		}
		return &d, hasDeviceBeenCreated, nil
	}

	if err := r.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Omit("id", "external_id", "is_deleted", "created_at", "updated_at", "deleted_at").
		Create(device).Error; err != nil {
		return nil, hasDeviceBeenCreated, err
	}
	hasDeviceBeenCreated = true

	return device, true, nil
}