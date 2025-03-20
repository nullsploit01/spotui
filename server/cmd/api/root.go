package main

import (
	"fmt"
	"log/slog"
)

type application struct {
	logger *slog.Logger
	config AppConfig
}

func Run(logger *slog.Logger) error {
	config, err := GetAppConfig()
	if err != nil {
		return fmt.Errorf("failed to get app config: %w", err)
	}

	app := &application{
		logger: logger,
		config: config,
	}

	app.logger.Info("starting server")
	return app.serveHTTP()
}
