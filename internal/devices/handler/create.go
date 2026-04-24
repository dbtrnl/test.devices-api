package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Brand string `json:"brand" binding:"required,min=2"`
	State string `json:"state"`
}

func (h *DeviceHandler) Create(c *gin.Context) {
	var req createDeviceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request", // Change for a constant
			"details": err.Error(),
		})
		return
	}

	input, err := dto.NewCreateDeviceInput(req.Name, req.Brand, req.State)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	created, err := h.svc.Create(c.Request.Context(), input)
	if err != nil {
		if domain.IsErrorCode(err, domain.ErrDeviceExistsDeleted) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create device",
		})
		return
	}

	response := dto.ToDeviceResponse(created)
	c.JSON(http.StatusCreated, response)
}
