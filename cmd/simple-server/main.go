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

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"simple-server/server"
	"simple-server/utils"
)

func main() {

	log := log.New(
		os.Stdout,
		"SIMPLE_SERVICE: ",
		log.LstdFlags|log.Lmicroseconds|log.Lshortfile,
	)

	logZap := utils.DevelopmentLogger()

	if err := run(log, logZap); err != nil {
		logZap.Fatal(err.Error())
	}

}

func run(log *log.Logger, logZap *zap.Logger) error {

	config, err := utils.LoadConfig("./")
	if err != nil {
		return fmt.Errorf("cannot load config: %v", err)
	}

	if config.Environment == "production" {
		logZap = utils.ProductionLogger()
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	srv := server.NewServer(router, logZap, shutdown)
	srv.SetRoutes()

	logZap.Sugar().Infof("Starting server on %s", config.ServerAddress)
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
		logZap.Sugar().Infof("shutdown started with sig: %v", sig)

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
