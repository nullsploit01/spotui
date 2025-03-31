package auth

import (
	"fmt"

	"github.com/nullsploit01/spotui/cmd/models"
	"github.com/nullsploit01/spotui/cmd/utils"
)

func StartAuthServer(clientID, redirectURI string) (models.AuthResponse, error) {
	verifier, err := utils.GenerateCodeVerifier()
	if err != nil {
		return models.AuthResponse{}, err
	}

	challenge := utils.GenerateCodeChallenge(verifier)
	state, err := utils.GenerateRandomString(10)
	if err != nil {
		return models.AuthResponse{}, err
	}

	authURL := GetAuthURL(clientID, redirectURI, challenge, state)

	fmt.Println("Opening browser for authentication...")
	utils.OpenBrowser(authURL)

	code, err := StartRedirectServer()
	if err != nil {
		return models.AuthResponse{}, err
	}

	token, err := ExchangeCodeForToken(code, verifier, clientID, redirectURI)
	if err != nil {
		return models.AuthResponse{}, err
	}

	return token, err
}
