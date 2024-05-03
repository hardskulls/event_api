package app

import (
	"log/slog"
	"os"
)

func InitLogger() (*slog.Logger, error) {
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewTextHandler(os.Stdout, opts)
	log := slog.New(handler)

	return log, nil
}
