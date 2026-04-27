package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/request"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

type listDevicesQuery struct {
	Brand *string `form:"brand"`
	State *string `form:"state"`
}

func (h *DeviceHandler) List(c *gin.Context) {
	var req listDevicesQuery

	if !request.BindQuery(c, &req) {
		return
	}

	input, err := dto.NewListDeviceInput(req.Brand, req.State)
	if !response.HandleValidationError(c, err) {
		return
	}

	devices, err := h.svc.List(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err, listDevicesErrMap)
		return
	}

	c.JSON(http.StatusOK, devices)
}
