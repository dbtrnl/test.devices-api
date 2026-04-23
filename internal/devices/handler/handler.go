package handler

import (
	"github.com/dbtrnl/test.devices-api/internal/devices/usecase"
)

type DeviceHandler struct {
	svc *usecase.DeviceService
}

func NewDeviceHandler(svc *usecase.DeviceService) *DeviceHandler {
	return &DeviceHandler{svc: svc}
}
