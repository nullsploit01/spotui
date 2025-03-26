package auth

import (
	"fmt"

	"github.com/nullsploit01/spotui/cmd/utils"
)

func StartAuthServer(clientID, redirectURI string) error {
	verifier, err := utils.GenerateCodeVerifier()
	if err != nil {
		return err
	}

	challenge := utils.GenerateCodeChallenge(verifier)
	state, err := utils.GenerateRandomString(10)
	if err != nil {
		return err
	}

	authURL := GetAuthURL(clientID, redirectURI, challenge, state)

	fmt.Println("Opening browser for authentication...")
	utils.OpenBrowser(authURL)

	code, err := StartRedirectServer()
	if err != nil {
		return err
	}

	token, err := ExchangeCodeForToken(code, verifier, clientID, redirectURI)
	if err != nil {
		return err
	}

	fmt.Println("Access Token Response:")
	utils.PrintPrettyPrintJSON(token)
	return nil
}
