package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dbtrnl/test.devices-api/internal/infra/logger"
	"github.com/dbtrnl/test.devices-api/internal/infra/server"
)

func main() {
	logger.Init()
	logger.Info("starting server...")

	srv := server.New()

	go func() {
		if err := srv.Start(); err != nil {
			logger.Error(fmt.Sprintf("server error: %v", err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("forcing server shutdown due to error: %v", err)
	}

	logger.Info("server gracefully stopped.")
}