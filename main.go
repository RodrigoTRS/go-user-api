package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	"user-api/src/api"
	"user-api/src/db"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("all systems offline")
}

const PORT = 3333

func run() error {
	db := db.Create()

	r := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      r,
	}

	slog.Info(fmt.Sprintf("server running on port: %d", PORT))

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
