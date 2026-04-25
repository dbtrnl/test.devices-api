package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/gin-gonic/gin"
)

var (
	deleteDeviceErrMap = map[domain.ErrorCode]func(*gin.Context, error){
		domain.ErrDeviceNotFoundCode: func(c *gin.Context, err error) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceInUse: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceAlreadyDeleted: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
	}
	createDeviceErrMap = map[domain.ErrorCode]func(*gin.Context, error){
		domain.ErrDeviceExistsDeleted: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
	}
	getByExternalIDErrMap = map[domain.ErrorCode]func(*gin.Context, error){}
	listDevicesErrMap     = map[domain.ErrorCode]func(*gin.Context, error){}
)
