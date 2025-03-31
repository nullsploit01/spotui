package cmd

import (
	"github.com/nullsploit01/spotui/cmd/auth"
	"github.com/nullsploit01/spotui/cmd/ui"
)

type App struct {
	config      *Config
	accessToken string
}

func start() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	token, err := auth.GetAccessToken(config.ClientID, config.RedirectURI)
	if err != nil {
		return err
	}

	app := &App{
		config:      config,
		accessToken: token,
	}

	return ui.StartUI()
}
