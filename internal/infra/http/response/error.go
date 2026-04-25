package response

import (
	"errors"
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, m map[domain.ErrorCode]func(*gin.Context, error)) {
	if appErr, ok := errors.AsType[*domain.AppError](err); ok {
		if handler := m[appErr.Code]; handler != nil {
			handler(c, err)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}

func HandleValidationError(c *gin.Context, err error) bool {
	if err == nil {
		return true
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	return false
}
