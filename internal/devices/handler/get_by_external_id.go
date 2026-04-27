package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

func (h *DeviceHandler) GetByExternalID(c *gin.Context) {
	externalID := c.Param("external_id")

	input, err := dto.NewDeleteDeviceInput(externalID)
	if !response.HandleValidationError(c, err) {
		return
	}

	device, err := h.svc.GetByExternalID(c, input)
	if err != nil {
		response.HandleError(c, err, getByExternalIDErrMap)
		return
	}
	response := dto.ToDeviceResponse(device)

	c.JSON(http.StatusOK, response)
}
