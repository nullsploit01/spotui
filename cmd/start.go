package cmd

import (
	"github.com/nullsploit01/spotui/cmd/auth"
	"github.com/nullsploit01/spotui/cmd/ui"
)

type App struct {
	config *Config
}

func start() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	err = auth.StartAuthServer(config.ClientID, config.RedirectURI)
	if err != nil {
		return err
	}

	return ui.StartUI()
}
