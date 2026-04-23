package config

import (
	"fmt"
	"os"

	"github.com/dbtrnl/test.devices-api/pkg/logger"
	"github.com/joho/godotenv"
)

const (
	EnvLocalFilename = ".env.local"
	EnvProdFilename  = ".env.prod"
)

type EnvConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	IsLocal    bool
}

func LoadConfig() (EnvConfig, error) {
	env, isSet := os.LookupEnv("ENV")
	if !isSet {
		return EnvConfig{}, fmt.Errorf("ENV variable not set")
	}
	if env == "" {
		env = "local"
	}

	if err := loadEnvFile(env); err != nil {
		logger.Info("no env file found, using default local environment variables...")
	}

	defaultEnv := EnvConfig{
		DBHost:     mustGet("DB_HOST"),
		DBPort:     mustGet("DB_PORT"),
		DBUser:     mustGet("DB_USER"),
		DBPassword: mustGet("DB_PASSWORD"),
		DBName:     mustGet("DB_NAME"),
		Port:       mustGet("PORT"),
		IsLocal:    env == "local",
	}

	return defaultEnv, nil
}

func loadEnvFile(env string) error {
	var filename string

	switch env {
	case "prod":
		filename = ".env.prod"
	default:
		filename = ".env.local"
	}

	if err := godotenv.Load(filename); err != nil {
		msg := fmt.Sprintf("could not load %s: %v", filename, err)
		logger.Info(msg)

		return fmt.Errorf(msg)
	}

	return nil
}

func mustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("missing environment variable: %s", key))
	}
	return val
}
