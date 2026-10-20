package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/ghafbu/padegan-backend/app/http"

	"github.com/ghafbu/padegan-backend/lib/bootstrap"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiilogger/myzap"
)

func main() {
	// Init Env
	thisEnv := os.Getenv("ENV_PATH")

	thisWd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current working directory: %v", err)
	}

	// Initialize the application using the configuration path based on the working directory and environment mode
	thisApp, err := bootstrap.Start(filepath.Join(thisWd, thisEnv))
	if err != nil {
		log.Fatalf("Application failed:%v", err)
	}

	//Log the loaded application configuration for debugging
	log.Printf("Configuration: %#v", thisApp.Config)

	// Setup signal handling to allow graceful shutdown on receiving termination signals
	thisQuit := make(chan os.Signal, 1)

	// Listen for SIGINT or similar termination signals, more SIGX (SIGINT, SIGTERM, etc)
	signal.Notify(thisQuit, os.Interrupt)

	// *********************** Start Servers **************************

	// HTTP Server
	hs := http.New(thisApp)

	go func() {
		if err = hs.Init(); err != nil {
			thisApp.Logger.Panicf("Server.HTTP.Run: %v", err)
		}
	}()

	// GRPC Server
	go func() {}()

	<-thisQuit // Block execution until a termination signal is received

	/*
	*
	*
	* Remove Memory Allocate Block
	*
	*
	 */
	// Create a context with a timeout for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		thisApp.Config.Other.CMDGracefulShutDownTimeOut,
	)
	defer shutdownCancel()

	// Attempt to gracefully shut down the HTTP server
	if err = hs.ShutdownWithContext(shutdownCtx); err != nil {
		thisApp.Logger.Errorf(
			myzap.HTTPServer.String(),
			fmt.Sprintf("%s: %v", "failed to gracefully shutdown", err),
		)
	}

	// Close PostgreSQL connections gracefully
	if err = thisApp.Postgres.Close(); err != nil {
		thisApp.Logger.Errorf(
			myzap.Postgres.String(),
			fmt.Sprintf("%s: %v", "failed to close PostgreSQL connection", err),
		)
	}

	// Close Redis connections gracefully
	//if err = thisApp.Redis.Close(); err != nil {
	//	thisApp.Logger.Errorf(
	//		myzap.Redis.String(),
	//		fmt.Sprintf("%s: %v", "failed to close Redis connection:", err),
	//	)
	//}

	// Wait for the timeout context to complete before exiting
	<-shutdownCtx.Done()
}
