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

	err := r.db.WithContext(ctx).
		First(&device, "external_id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.NewDeviceNotFoundError(id)
	}
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *deviceRepository) Create(ctx context.Context, device *domain.Device) (*domain.Device, error) {
	var d domain.Device

	err := r.db.WithContext(ctx).
		// Omit("id").
		Where("name = ? AND brand = ?", device.Name, device.Brand).
		First(&d).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		if d.IsDeleted {
			return nil, domain.NewErrDeviceAExistsDeletedError(d.Name, d.Brand, d.ExternalID)
		}
		return &d, nil
	}

	if err := r.db.WithContext(ctx).
		Omit("id", "external_id", "is_deleted", "created_at", "updated_at", "deleted_at").
		Create(device).Error; err != nil {
		return nil, err
	}

	return device, nil
}

// func (r *deviceRepository) List(ctx context.Context, filter domain.DeviceFilter) ([]domain.Device, error) {
//     var devices []domain.Device
//     db := r.db.WithContext(ctx).Model(&domain.Device{})

//     if filter.Brand != nil {
//         db = db.Where("brand = ?", *filter.Brand)
//     }
//     if filter.State != nil {
//         db = db.Where("state = ?", *filter.State)
//     }
//     if err := db.Find(&devices).Error; err != nil {
//         return nil, err
//     }
//     return devices, nil
// }
