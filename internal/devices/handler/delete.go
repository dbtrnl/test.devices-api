package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

func (h *DeviceHandler) DeleteByExternalID(c *gin.Context) {
	externalID := c.Param("external_id")

	input, err := dto.NewDeleteDeviceInput(externalID)
	if !response.HandleValidationError(c, err) {
		return
	}

	err = h.svc.DeleteByExternalID(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err, deleteDeviceErrMap)
		return
	}

	c.Status(http.StatusNoContent)
}
