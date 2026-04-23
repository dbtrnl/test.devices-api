package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dbtrnl/test.devices-api/internal/infra/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	r := gin.New()
	r.Use(gin.Recovery())

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %v", err))
	}

	s := &Server{}
	s.registerRoutes(r)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	return s
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}


