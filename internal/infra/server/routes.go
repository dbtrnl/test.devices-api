package server

import (
	"fmt"

	"github.com/dbtrnl/test.devices-api/internal/devices/handler"
	"github.com/dbtrnl/test.devices-api/internal/devices/repository"
	"github.com/dbtrnl/test.devices-api/internal/devices/usecase"
	"github.com/dbtrnl/test.devices-api/internal/infra/buildinfo"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes(r *gin.Engine) {
	r.GET("/health", s.health)

	dbConn, err := s.GetDbConn()
	if err != nil {
		panic(fmt.Errorf("failed to get database connection: %v", err))
	}

	repo := repository.NewDeviceRepository(dbConn)
	uc := usecase.NewDeviceService(repo)
	h := handler.NewDeviceHandler(uc)

	r.GET("/devices/:id", h.GetDevice)
}

func (s *Server) health(c *gin.Context) {
    c.JSON(200, gin.H{
        "status":    "ok",
        "version":   buildinfo.Version,
        "commit":    buildinfo.Commit,
        "buildTime": buildinfo.BuildTime,
    })
}
