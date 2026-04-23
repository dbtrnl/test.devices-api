package domain

const (
	DeviceStateActive   DeviceState = "active"
	DeviceStateInactive DeviceState = "inactive"
	DeviceStateInUse    DeviceState = "in_use"
)

type Device struct {
	ID         string      `json:"id"`
	ExternalID string      `json:"external_id"`
	Name       string      `json:"name"`
	Brand      string      `json:"brand"`
	State      DeviceState `json:"state"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at,omitzero"`
	DeletedAt  string      `json:"deleted_at,omitzero"`
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
