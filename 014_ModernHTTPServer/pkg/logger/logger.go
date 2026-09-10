package logger

import (
	"log/slog"
	"os"
)

// New creates a new structured JSON logger writing to standard output.
func New() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
}
