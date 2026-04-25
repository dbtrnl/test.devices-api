package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/gin-gonic/gin"
)

type listDevicesQuery struct {
	Brand *string `json:"brand"`
	State *string `json:"state"`
}

func (h *DeviceHandler) List(c *gin.Context) {
	var req listDevicesQuery

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid query params",
		})
		return
	}

	input, err := dto.NewListDeviceInput(req.Brand, req.State)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	devices, err := h.svc.List(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list devices",
		})
		return
	}

	c.JSON(http.StatusOK, devices)
}
