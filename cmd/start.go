package cmd

import (
	"github.com/nullsploit01/spotui/cmd/auth"
	"github.com/nullsploit01/spotui/cmd/ui"
	store "github.com/nullsploit01/spotui/cmd/utils/store"
)

type App struct {
	config *Config
}

func start() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	token, err := store.Get("refresh_token")
	if err != nil {
		return err
	}

	if token == "" {
		err = auth.StartAuthServer(config.ClientID, config.RedirectURI)
		if err != nil {
			return err
		}
	}

	return ui.StartUI()
}
