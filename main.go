package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github-watcher/config"
	"github-watcher/routes"
)

func main() {
	// Initialize configuration
	cfg := &config.Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}

	// Initialize database
	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize router
	router := routes.NewRouter(db)
	router.SetupRoutes()

	// Get the Fiber app
	app := router.GetApp()

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":" + cfg.ServerPort); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.ServerPort)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited properly")
}
