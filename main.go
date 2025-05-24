package main

import (
	"context"
	"fmt"
	"log"
	mongodb "lumel/internal/database"
	"lumel/internal/dataloader"
	"lumel/internal/router"
	"lumel/pkg/logger"
	cfg "lumel/pkg/settings"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Initialize application configuration from environment variables and settings files
	config, err := initializeConfig()
	if err != nil {
		fmt.Println("Not able to get config files")
	}

	// Initialize the logger with settings from the configuration
	initializeLogger(config)
	logger.Log.Info("Logger Initialized")

	// Connect to MongoDB
	mongoClient, err := mongodb.NewClient(cfg.Config)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a context for cleanup operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ensure MongoDB connection is closed on exit
	defer func() {
		if err := mongoClient.Close(ctx); err != nil {
			log.Printf("Error closing MongoDB connection: %v", err)
		}
	}()

	logger.Log.Info("MongoDB Initialized")

	// Initialize the HTTP router with the registered routes
	router := router.GetRouter()
	logger.Log.Info("Router Initialized")

	// load data using cron job on every day
	c := cron.New()
	go func() {
		c.AddFunc("0 0 * * *", func() {
			go dataloader.LoadSalesData()
		}) // Runs every day at midnight
		c.Start()
	}()

	// Signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Start the HTTP server on the configured port, log and handle any errors
		if err := router.Listen(fmt.Sprintf(":%s", config.AppPort)); err != nil {
			logger.Log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	// Wait for a termination signal before gracefully closing services
	<-quit

	logger.Log.Info("Shutting down server...")

	logger.Log.Info("Server gracefully stopped.")
}

/*
initializeConfig loads the application configuration from environment variables
and other settings files using the InitConfig function from the settings package.
*/
func initializeConfig() (cfg.Configuration, error) {
	return cfg.InitConfig()
}

/*
initializeLogger sets up the logging system using the settings specified
in the configuration file (e.g., log file name, size, retention, and level).
*/
func initializeLogger(conf cfg.Configuration) {
	logger.InitLogger(
		conf.Logger.FileName,
		conf.Logger.FileSize,
		conf.Logger.MaxLogFile,
		conf.Logger.MaxRetention,
		conf.Logger.CompressLog,
		conf.Logger.Level,
	)
}
