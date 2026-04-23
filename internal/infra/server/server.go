package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/infra/deps"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func New(c *deps.Container) (*Server, error) {
	r := gin.New()
	r.Use(gin.Recovery())

	if err := SetupRoutes(r, c); err != nil {
		return nil, err
	}

	s := &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%s", c.Config.Port),
			Handler: r,
		},
	}

	return s, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
