package auth

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/nullsploit01/spotui/cmd/models"
	store "github.com/nullsploit01/spotui/cmd/utils/store"
)

func GetAccessToken(clientID, redirectURI string) (string, error) {
	if refreshToken, err := store.Get("refresh_token"); err == nil && refreshToken != "" {
		if tokenResp, err := RefreshToken(refreshToken, clientID); err == nil {
			return tokenResp.AccessToken, nil
		}
	}

	authResp, err := StartAuthServer(clientID, redirectURI)
	if err != nil {
		return "", err
	}

	if authResp.RefreshToken != "" {
		_ = store.Set("refresh_token", authResp.RefreshToken)
	}

	return authResp.AccessToken, nil
}

func RefreshToken(refreshToken, clientId string) (models.AuthResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientId)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", data)
	if err != nil {
		return models.AuthResponse{}, err
	}
	defer resp.Body.Close()

	var result models.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
