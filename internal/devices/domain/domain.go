package domain

import "time"

const (
	DeviceStateActive   DeviceState = "active"
	DeviceStateInactive DeviceState = "inactive"
	DeviceStateInUse    DeviceState = "in_use"
)

type Device struct {
	ID         uint64
	ExternalID string
	Name       string
	Brand      string
	State      DeviceState
	IsDeleted  bool
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
type DeviceState string

func (Device) TableName() string {
	return "devices"
}

func (d Device) IsInactive() bool {
	return d.State == DeviceStateInactive
}

func (d Device) IsInUse() bool {
	return d.State == DeviceStateInUse
}

type DeviceFilter struct {
	Brand *string
	State *string

	Limit  *int
	Offset *int
}
