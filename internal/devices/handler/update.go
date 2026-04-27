package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/request"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

type updateDeviceRequest struct {
	Name  *string `json:"name" binding:"omitempty,min=2"`
	Brand *string `json:"brand" binding:"omitempty,min=2"`
	State *string `json:"state"`
}

func (h *DeviceHandler) Update(c *gin.Context) {
	externalID := c.Param("external_id")

	var req updateDeviceRequest
	if !request.BindJSON(c, &req) {
		return
	}

	input, err := dto.NewUpdateDeviceInput(externalID, req.Brand, req.Name, req.State)
	if !response.HandleValidationError(c, err) {
		return
	}

	device, err := h.svc.Update(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err, updateDeviceErrMap)
		return
	}

	c.JSON(http.StatusOK, dto.ToDeviceResponse(device))
}
