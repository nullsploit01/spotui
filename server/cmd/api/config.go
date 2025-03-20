package main

import "github.com/nullsploit01/spotui-server/internal/env"

type AppConfig struct {
	httpPort int
	debug    bool
}

func GetAppConfig() (AppConfig, error) {
	return AppConfig{
		httpPort: env.GetInt("PORT", 8000),
		debug:    env.GetBool("DEBUG", false),
	}, nil
}
