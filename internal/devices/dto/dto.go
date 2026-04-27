package dto

import (
	"fmt"
	"time"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/google/uuid"
)

type CreateDeviceInput struct {
	Name  string
	Brand string
	State domain.DeviceState
}

type DeviceUUIDInput struct {
	ExternalID string
}

type ListDevicesInput struct {
	Brand *string
	State *domain.DeviceState
}

type UpdateDeviceInput struct {
	ExternalID string
	Name       *string
	Brand      *string
	State      *domain.DeviceState
}

type DeviceResponse struct {
	ID         uint64  `json:"id"`
	ExternalID string  `json:"external_id"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	State      string  `json:"state"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
	DeletedAt  *string `json:"deleted_at,omitempty"`
	IsDeleted  bool    `json:"is_deleted,omitempty"`
}

func NewCreateDeviceInput(name, brand string, state string) (CreateDeviceInput, error) {
	var input CreateDeviceInput
	if len(name) < 2 {
		return CreateDeviceInput{}, fmt.Errorf("name must be at least 2 characters")
	}
	if len(brand) < 2 {
		return CreateDeviceInput{}, fmt.Errorf("brand must be at least 2 characters")
	}

	input.Name = name
	input.Brand = brand

	if state != "" {
		if !domain.IsValidDeviceState(state) {
			return CreateDeviceInput{}, fmt.Errorf("invalid state: %s", state)
		}

		s := domain.DeviceState(state)
		input.State = s
	}

	return input, nil
}

func NewListDeviceInput(brand, state *string) (ListDevicesInput, error) {
	var input ListDevicesInput

	if brand != nil {
		input.Brand = brand
	}
	if state != nil {
		if !domain.IsValidDeviceState(*state) {
			return ListDevicesInput{}, fmt.Errorf("invalid state: %s", *state)
		}

		s := domain.DeviceState(*state)
		input.State = &s
	}

	return input, nil
}

func NewDeleteDeviceInput(uuidStr string) (DeviceUUIDInput, error) {
	err := uuid.Validate(uuidStr)
	if err != nil {
		return DeviceUUIDInput{}, fmt.Errorf("invalid uuid: %s", uuidStr)
	}

	return DeviceUUIDInput{
		ExternalID: uuidStr,
	}, nil
}

func NewUpdateDeviceInput(externalID string, name, brand, state *string) (UpdateDeviceInput, error) {
	err := uuid.Validate(externalID)
	if err != nil {
		return UpdateDeviceInput{}, fmt.Errorf("invalid uuid: %s", externalID)
	}

	input := UpdateDeviceInput{
		ExternalID: externalID,
	}

	if name != nil {
		if len(*name) < 2 {
			return UpdateDeviceInput{}, fmt.Errorf("name must be at least 2 characters")
		}
		input.Name = name
	}

	if brand != nil {
		if len(*brand) < 2 {
			return UpdateDeviceInput{}, fmt.Errorf("brand must be at least 2 characters")
		}
		input.Brand = brand
	}

	if state != nil {
		if !domain.IsValidDeviceState(*state) {
			return UpdateDeviceInput{}, fmt.Errorf("invalid state: %s", *state)
		}
		s := domain.DeviceState(*state)
		input.State = &s
	}

	return input, nil
}

func ToDeviceResponse(d *domain.Device) DeviceResponse {
	var updatedAt, deletedAt *string
	if d.UpdatedAt != nil {
		s := d.UpdatedAt.UTC().Format(time.RFC3339)
		updatedAt = &s
	}

	if d.DeletedAt != nil {
		del := d.DeletedAt.UTC().Format(time.RFC3339)
		deletedAt = &del
	}

	return DeviceResponse{
		ID:         d.ID,
		ExternalID: d.ExternalID,
		Name:       d.Name,
		Brand:      d.Brand,
		State:      string(d.State),
		CreatedAt:  d.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
	}
}
