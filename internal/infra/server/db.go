package server

import (
	"fmt"
	"sync"

	"github.com/dbtrnl/test.devices-api/internal/infra/config"
	"github.com/dbtrnl/test.devices-api/internal/infra/logger"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func (s *Server) GetDbConn() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dbInstance, err = s.initDbConn()
	})
	return dbInstance, err
}

func (s *Server) initDbConn() (*gorm.DB, error) {
	var sslDisable string

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// SSL disabled on local env since this is a take-home coding exercise. Also used to make tests easier.
	// // This obviously would never be used in production.
	if cfg.IsLocal {
		sslDisable = "?sslmode=disable"
	}

	logger.Debug(fmt.Sprintf("using cfg struct: %+v", cfg))
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, sslDisable,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	logger.Info("database connection established")

	return db, nil
}
