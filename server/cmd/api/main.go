package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	if err := Run(logger); err != nil {
		trace := debug.Stack()
		logger.Error(err.Error(), "trace", string(trace))
		os.Exit(1)
	}
}
