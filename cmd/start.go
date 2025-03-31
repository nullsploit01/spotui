package cmd

import (
	"github.com/nullsploit01/spotui/cmd/auth"
	"github.com/nullsploit01/spotui/cmd/ui"
)

type App struct {
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

	return ui.StartUI(token)
}
