package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Level      string `mapstructure:"level" default:"info"`
	Format     string `mapstructure:"format" default:"json"` // json or console
	TimeFormat string `mapstructure:"time_format" default:"2006-01-02T15:04:05.000Z07:00"`
}

func New(cfg Config) zerolog.Logger {
	// Parse log level
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)

	// Set time field format
	if cfg.TimeFormat != "" {
		zerolog.TimeFieldFormat = cfg.TimeFormat
	} else {
		zerolog.TimeFieldFormat = time.RFC3339
	}

	var output io.Writer = os.Stdout

	// Configure output format
	if cfg.Format == "console" {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: cfg.TimeFormat,
		}
	}

	logger := zerolog.New(output).With().Timestamp().Caller().Logger()

	return logger
}

// NewDefault creates a logger with default settings (JSON format, info level)
func NewDefault() zerolog.Logger {
	return New(Config{
		Level:      "info",
		Format:     "json",
		TimeFormat: time.RFC3339,
	})
}

// NewDevelopment creates a logger optimized for development (console format, debug level)
func NewDevelopment() zerolog.Logger {
	return New(Config{
		Level:      "debug",
		Format:     "console",
		TimeFormat: "15:04:05",
	})
}
