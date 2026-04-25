package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/dto"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/request"
	"github.com/dbtrnl/test.devices-api/internal/infra/http/response"
	"github.com/gin-gonic/gin"
)

type createDeviceRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Brand string `json:"brand" binding:"required,min=2"`
	State string `json:"state"`
}

func (h *DeviceHandler) Create(c *gin.Context) {
	var req createDeviceRequest

	if !request.BindJSON(c, &req) {
		return
	}

	input, err := dto.NewCreateDeviceInput(req.Name, req.Brand, req.State)
	if !response.HandleValidationError(c, err) {
		return
	}

	created, err := h.svc.Create(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err, createDeviceErrMap)
		return
	}

	response := dto.ToDeviceResponse(created)
	c.JSON(http.StatusCreated, response)
}
