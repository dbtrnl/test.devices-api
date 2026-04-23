package repository

import (
	"context"
	"errors"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"gorm.io/gorm"
)

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *deviceRepository {
	return &deviceRepository{db: db}
}

func (r *deviceRepository) GetByExternalID(ctx context.Context, id string) (*domain.Device, error) {
	var device domain.Device

	err := r.db.WithContext(ctx).Debug().
		First(&device, "external_id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.NewDeviceNotFoundError(id)
	}
	if err != nil {
		return nil, err
	}

	return &device, nil
}
