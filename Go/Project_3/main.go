package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"project3/api"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to run application", "error", err)
		os.Exit(1)
		return
	}

	slog.Info("Application started")

}

func run() error {

	handler := api.NewHandler(os.Getenv("OMDB_API_KEY"))

	port := ":8080"

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         port,
		Handler:      handler,
	}

	slog.Info("Application started", "port", port)
	if err := s.ListenAndServe(); err != nil {

		return err
	}

	return errors.New("test error")
}
