package config

import (
	"github.com/arthben/go-payments-platform/services/payment/internal/infrastructure/db"
	"github.com/arthben/go-payments-platform/services/payment/internal/infrastructure/logger"
)

type Config struct {
	App      AppConfig       `mapstructure:"app"`
	Database db.Config       `mapstructure:"database"`
	Logger   logger.Config   `mapstructure:"logger"`
}

type AppConfig struct {
	Name        string `mapstructure:"name" default:"payment-service"`
	Environment string `mapstructure:"environment" default:"development"`
	Port        int    `mapstructure:"port" default:"8080"`
	GRPCPort    int    `mapstructure:"grpc_port" default:"9090"`
}

// NewDefault returns a configuration with default values
func NewDefault() *Config {
	return &Config{
		App: AppConfig{
			Name:        "payment-service",
			Environment: "development",
			Port:        8080,
			GRPCPort:    9090,
		},
		Database: db.Config{
			SSLMode:  "disable",
			Host:     "localhost",
			Port:     5432,
			Name:     "gpp_payment_service",
			User:     "postgres",
			Password: "12345",
			PoolSize: 10,
		},
		Logger: logger.Config{
			Level:      "debug",
			Format:     "console",
			TimeFormat: "15:04:05",
		},
	}
}
