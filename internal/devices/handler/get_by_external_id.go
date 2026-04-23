package handler

import "github.com/gin-gonic/gin"

func (h *DeviceHandler) GetByExternalID(c *gin.Context) {
	external_id := c.Param("external_id")

	device, err := h.svc.GetByExternalID(c, external_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, device)
}
