package config

import (
	"log/slog"
	"os"
)

func ConfigureLogger() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	logger := slog.New(handler)

	slog.SetDefault(logger)
}
