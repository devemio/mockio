package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devemio/mockio/internal/color"
	"github.com/devemio/mockio/internal/config"
	"github.com/devemio/mockio/internal/handler"
	"github.com/devemio/mockio/internal/logger"
	"github.com/devemio/mockio/internal/routing"
)

func main() {
	cfg := config.New()
	clr := color.New(cfg.IsOutputColored)
	log := logger.New(clr)

	routes, err := routing.Parse(cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Starting server at port %d...", cfg.Port)

	log.Debug("Creating routes")
	for _, route := range routes {
		log.Debug("%-6s %s", route.Method, route.Path)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler.New(log, clr, routes, cfg.ResponseLatency, cfg.Verbose),
	}

	go func() {
		if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Debug("Server started (version %s)", config.Version)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Debug("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Debug("Graceful shutdown complete")
}
