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

type DeleteDeviceInput struct {
	ExternalID string
}

type ListDevicesInput struct {
	Brand *string
	State *domain.DeviceState
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

func NewCreateDeviceInput(name, brand, state string) (CreateDeviceInput, error) {
	if len(name) < 2 {
		return CreateDeviceInput{}, fmt.Errorf("name must be at least 2 characters")
	}

	if len(brand) < 2 {
		return CreateDeviceInput{}, fmt.Errorf("brand must be at least 2 characters")
	}

	// default state
	deviceState := domain.DeviceState(state)
	if deviceState == "" {
		deviceState = domain.DeviceStateInactive
	}

	return CreateDeviceInput{
		Name:  name,
		Brand: brand,
		State: deviceState,
	}, nil
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

func NewDeleteDeviceInput(uuidStr string) (DeleteDeviceInput, error) {
	err := uuid.Validate(uuidStr)
	if err != nil {
		return DeleteDeviceInput{}, fmt.Errorf("invalid uuid: %s", uuidStr)
	}

	return DeleteDeviceInput{
		ExternalID: uuidStr,
	}, nil
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
