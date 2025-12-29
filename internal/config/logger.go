package config

import (
	"log/slog"
	"os"
)

func NewLogger() (*slog.Logger, error) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger, nil
}
