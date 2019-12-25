package main

import (
	"context"
	"flag"
	"fmt"
	"gophr/api/v1"
	"gophr/api/v2"
	"gophr/middleware"
	"gophr/view"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jayvib/golog"
)

var port string
var host string
var debug bool

const AppName = "Gophr"

func init() {
	flag.StringVar(&port, "port", "8080", "Port of the application")
	flag.StringVar(&host, "host", "127.0.0.1", "Host of the application")
	flag.BoolVar(&debug, "debug", false, "Debug mode")

	flag.Parse()

	if debug {
		golog.Warning("DEBUGGING MODE!")
		golog.SetLevel(golog.DebugLevel)
	}
}

func main() {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	view.RegisterHandlers(router)

	apiRouter := router.PathPrefix("/api").Subrouter()

	// API version 1
	apiRouterV1 := v1.RegisterHandlers(apiRouter)
	apiRouterV1.Use(middleware.AuthenticateMiddleware)

	// API version 2
	v2.RegisterHandlers(apiRouter)

	addr := fmt.Sprintf("%s:%s", host, port)

	// Graceful shutdown
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Server
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go gracefulShutdown(server, quit, done)

	golog.Info("Listening in ", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		log.Fatal(err)
	}

	<-done
	golog.Info("Server stopped")
}

func gracefulShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	golog.Warning("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		golog.Fatal("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
