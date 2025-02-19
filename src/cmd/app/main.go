package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vmlellis/port-sync/src/internal/domain/service"

	"github.com/vmlellis/port-sync/src/internal/adapters/config"
	"github.com/vmlellis/port-sync/src/internal/adapters/file"
	"github.com/vmlellis/port-sync/src/internal/adapters/storage"

	httpadapter "github.com/vmlellis/port-sync/src/internal/adapters/http"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Initialize storage and service
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	handler := httpadapter.NewPortHandler(portService)

	// Load initial data if enabled in config
	if config.LoadOnStartup {
		fmt.Println("Loading data from", config.PortsFile)
		if err := file.ProcessJSONFile(config.PortsFile, portService); err != nil {
			fmt.Println("Error loading data:", err)
			os.Exit(1)
		}
		fmt.Println("Data loaded successfully.")
	}

	// Create HTTP server
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	server := &http.Server{Addr: ":8080", Handler: mux}

	// Run the server in a goroutine
	go func() {
		fmt.Println("Server is running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting server:", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		fmt.Println("Error shutting down server:", err)
	}
	fmt.Println("Server gracefully stopped.")
}
