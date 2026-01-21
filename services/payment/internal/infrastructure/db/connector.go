package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type Config struct {
	SSLMode      string        `mapstructure:"sslmode"`
	Host         string        `mapstructure:"host"           default:"localhost"`
	Port         uint16        `mapstructure:"port"           default:"5432"`
	Name         string        `mapstructure:"name"           default:"gpp_payment_service"`
	User         string        `mapstructure:"user"           default:"postgres"`
	Password     string        `mapstructure:"password"       default:"12345"`
	PoolSize     int           `mapstructure:"pool_size"      default:"10"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"   default:"10s"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"  default:"10s"`
	EnableLogger bool          `mapstructure:"enable_logger"`
}

func New(ctx context.Context, wg *sync.WaitGroup, cfg Config, logger zerolog.Logger) (*pgxpool.Pool, error) {
	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolSize,
	)

	// Parse the connection string to create the config
	pgConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Configure pool settings
	pgConfig.MaxConns = int32(cfg.PoolSize)
	pgConfig.MinConns = 1
	pgConfig.MaxConnLifetime = time.Hour
	pgConfig.MaxConnIdleTime = 30 * time.Minute
	pgConfig.HealthCheckPeriod = 1 * time.Minute

	// Attach logger to pgx if enabled
	if cfg.EnableLogger {
		pgConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger: tracelog.LoggerFunc(func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
				logEvent := logger.WithLevel(zerologLevelFromPgx(level))
				for k, v := range data {
					logEvent = logEvent.Interface(k, v)
				}
				logEvent.Msg(msg)
			}),
			LogLevel: tracelog.LogLevelDebug,
		}
	}

	// Create the pool
	pool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create db connection pool: %w", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping db: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		pool.Close()
	}()

	return pool, nil
}

// zerologLevelFromPgx converts pgx log levels to zerolog levels
func zerologLevelFromPgx(level tracelog.LogLevel) zerolog.Level {
	switch level {
	case tracelog.LogLevelTrace, tracelog.LogLevelDebug:
		return zerolog.DebugLevel
	case tracelog.LogLevelInfo:
		return zerolog.InfoLevel
	case tracelog.LogLevelWarn:
		return zerolog.WarnLevel
	case tracelog.LogLevelError:
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
