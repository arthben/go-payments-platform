package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arthben/go-payments-platform/services/payment/internal/config"
	"github.com/arthben/go-payments-platform/services/payment/internal/infrastructure/db"
	"github.com/arthben/go-payments-platform/services/payment/internal/infrastructure/logger"
)

func main() {
	// Load configuration
	cfg := config.NewDefault()

	// Initialize logger
	log := logger.New(cfg.Logger)
	log.Info().
		Str("service", cfg.App.Name).
		Str("environment", cfg.App.Environment).
		Msg("Starting payment service")

	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create wait group for graceful shutdown
	var wg sync.WaitGroup

	// Initialize database connection
	log.Info().Msg("Connecting to database")
	dbPool, err := db.New(ctx, &wg, cfg.Database, log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer dbPool.Close()

	log.Info().Msg("Database connection established")

	// TODO: Initialize repositories, services, and handlers

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutting down payment service")

	cancel()
	wg.Wait()

	log.Info().Msg("Payment service stopped")
	fmt.Println("Payment service stopped successfully")
}
