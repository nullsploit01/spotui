package cmd

import (
	"fmt"

	"github.com/nullsploit01/spotui/cmd/auth"
	"github.com/nullsploit01/spotui/cmd/utils"
)

type App struct {
	config *Config
}

func start() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	app := &App{
		config: config,
	}

	verifier, err := utils.GenerateCodeVerifier()
	if err != nil {
		return err
	}

	challenge := utils.GenerateCodeChallenge(verifier)
	state, err := utils.GenerateRandomString(10)
	if err != nil {
		return err
	}

	authURL := auth.GetAuthURL(app.config.ClientID, app.config.RedirectURI, challenge, state)

	fmt.Println("Opening browser for authentication...")
	utils.OpenBrowser(authURL)

	code, err := auth.StartRedirectServer()
	if err != nil {
		return err
	}

	token, err := auth.ExchangeCodeForToken(code, verifier, app.config.ClientID, app.config.RedirectURI)
	if err != nil {
		return err
	}

	fmt.Println("Access Token Response:")
	utils.PrintPrettyPrintJSON(token)
	return nil
}
