package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"powerfactors/assignment/internal/api"
	"powerfactors/assignment/internal/server"
	"powerfactors/assignment/internal/service"
	"syscall"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()

	// arguments check
	args := os.Args[1:]
	if len(args) < 2 {
		logger.Fatal("address and port arguments are required")
	}

	// initializations
	router := mux.NewRouter()
	httpServer := &http.Server{
		Addr: args[0] + `:` + args[1],
	}

	taskService := service.NewTaskService(logger)
	timestampHandler := api.NewTimestampHandler(taskService)
	server := server.NewListener(router, logger, httpServer, timestampHandler)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	server.Route()
	go server.Start()

	// graceful shutdown
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("Listener shutdown failed: %v", zap.Error(err))
	}

	logger.Info("Service gracefully stopped")

}
