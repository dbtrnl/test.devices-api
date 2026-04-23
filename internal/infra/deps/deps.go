package deps

import (
	"fmt"
	"sync"

	"github.com/dbtrnl/test.devices-api/pkg/config"
	"github.com/dbtrnl/test.devices-api/pkg/logger"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Container struct {
	Config *config.EnvConfig
	DB     *gorm.DB
}

var (
	instance *Container
	once     sync.Once
	initErr  error
)

func NewContainer() (*Container, error) {
	once.Do(func() {
		instance, initErr = newContainerInstance()
	})
	return instance, initErr
}

func newContainerInstance() (*Container, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	db, err := initDbConn(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return &Container{
		Config: &cfg,
		DB:     db,
	}, nil
}

func initDbConn(cfg *config.EnvConfig) (*gorm.DB, error) {
	var sslDisable string

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
