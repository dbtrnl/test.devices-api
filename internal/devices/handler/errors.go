package handler

import (
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/gin-gonic/gin"
)

var (
	createDeviceErrMap = map[domain.ErrorCode]func(*gin.Context, error){
		domain.ErrDeviceExistsDeleted: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
	}
	deleteDeviceErrMap = map[domain.ErrorCode]func(*gin.Context, error){
		domain.ErrDeviceNotFoundCode: func(c *gin.Context, err error) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceInUse: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceDeleted: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
	}
	getByExternalIDErrMap = map[domain.ErrorCode]func(*gin.Context, error){}
	listDevicesErrMap     = map[domain.ErrorCode]func(*gin.Context, error){}
	updateDeviceErrMap    = map[domain.ErrorCode]func(*gin.Context, error){
		domain.ErrDeviceInUse: func(c *gin.Context, err error) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceNotFoundCode: func(c *gin.Context, err error) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		},
		domain.ErrDeviceDeleted: func(c *gin.Context, err error) {
			c.JSON(http.StatusGone, gin.H{"error": err.Error()})
		},
	}
)
