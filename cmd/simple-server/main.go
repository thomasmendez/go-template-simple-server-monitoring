package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simple-server/server"
	"simple-server/util"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	log := log.New(
		os.Stdout,
		"SIMPLE_SERVICE: ",
		log.LstdFlags|log.Lmicroseconds|log.Lshortfile,
	)
	if err := run(log); err != nil {
		log.Println("main: error: ", err)
		os.Exit(1)
	}

}

func run(log *log.Logger) error {

	// Configurations
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Start Service
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	srv := server.NewServer(router, log, shutdown)
	srv.SetRoutes()

	log.Printf("Starting server on %s", config.ServerAddress)
	api := http.Server{
		Addr:     config.ServerAddress,
		Handler:  srv,
		ErrorLog: log,
	}

	// Make a chanel to listen for errors coming from the listner. Use a
	// buffered channel so the goroutine can exit if we don't collect this error
	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown
	// Blocking main and waiting for shutdown
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		log.Printf("shutdown started with sig: %v", sig)
		defer log.Println()

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Asking listener to shutdown and shed load
		if err := api.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}
	return nil
}
