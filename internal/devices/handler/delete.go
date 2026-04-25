package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/gin-gonic/gin"
)

func (h *DeviceHandler) DeleteByExternalID(c *gin.Context) {
	externalID := c.Param("external_id")

	input, err := dto.NewDeleteDeviceInput(externalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.svc.DeleteByExternalID(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete device",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
