package main

import "log/slog"

type application struct {
	logger *slog.Logger
}

func Run(logger *slog.Logger) error {
	app := &application{
		logger: logger,
	}

	app.logger.Info("starting server")
	return nil
}
