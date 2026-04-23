package server

import (
	"github.com/dbtrnl/test.devices-api/internal/devices/handler"
	"github.com/dbtrnl/test.devices-api/internal/devices/repository"
	"github.com/dbtrnl/test.devices-api/internal/devices/usecase"
	"github.com/dbtrnl/test.devices-api/internal/infra/buildinfo"
	"github.com/dbtrnl/test.devices-api/internal/infra/deps"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, c *deps.Container) error {
    r.GET("/health", health)

    repo := repository.NewDeviceRepository(c.DB)
    uc := usecase.NewDeviceService(repo)
    h := handler.NewDeviceHandler(uc)

    r.GET("/devices/:id", h.GetDevice)

    return nil
}

func health(c *gin.Context) {
    c.JSON(200, gin.H{
        "status":    "ok",
        "version":   buildinfo.Version,
        "commit":    buildinfo.Commit,
        "buildTime": buildinfo.BuildTime,
    })
}