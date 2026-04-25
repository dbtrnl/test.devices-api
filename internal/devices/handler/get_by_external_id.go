package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

func (h *DeviceHandler) GetByExternalID(c *gin.Context) {
	external_id := c.Param("external_id")

	device, err := h.svc.GetByExternalID(c, external_id)
	if err != nil {
		response.HandleError(c, err, getByExternalIDErrMap)
		return
	}
	c.JSON(http.StatusCreated, device)
}
