package dto

import (
	"fmt"
	"time"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
)

type CreateDeviceInput struct {
	Name  string
	Brand string
	State domain.DeviceState
}

type DeviceResponse struct {
	ID         uint64  `json:"id"`
	ExternalID string  `json:"external_id"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	State      string  `json:"state"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
	IsDeleted  bool    `json:"deleted_at,omitempty"`
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

func ToDeviceResponse(d *domain.Device) DeviceResponse {
	var updatedAt *string
	if d.UpdatedAt != nil {
		s := d.UpdatedAt.UTC().Format(time.RFC3339)
		updatedAt = &s
	}

	return DeviceResponse{
		ID:         d.ID,
		ExternalID: d.ExternalID,
		Name:       d.Name,
		Brand:      d.Brand,
		State:      string(d.State),
		CreatedAt:  d.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  updatedAt,
	}
}