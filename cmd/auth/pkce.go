package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nullsploit01/spotui/cmd/models"
)

var (
	scopes = "user-read-private user-read-email"
)

func GetAuthURL(clientID, redirectURI, challenge, state string) string {
	values := url.Values{}
	values.Set("client_id", clientID)
	values.Set("response_type", "code")
	values.Set("redirect_uri", redirectURI)
	values.Set("code_challenge_method", "S256")
	values.Set("code_challenge", challenge)
	values.Set("state", state)
	values.Set("scope", scopes)
	return fmt.Sprintf("https://accounts.spotify.com/authorize?%s", values.Encode())
}

func StartRedirectServer() (string, error) {
	var code string
	done := make(chan bool)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		fmt.Fprintf(w, "Authentication complete! You may now close this window.")
		done <- true
	})

	server := &http.Server{Addr: ":5555"}
	go func() {
		_ = server.ListenAndServe()
	}()

	<-done
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_ = server.Shutdown(ctxTimeout)
	return code, nil
}

func ExchangeCodeForToken(code, verifier, clientID, redirectURI string) (models.AuthResponse, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("code_verifier", verifier)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", data)
	if err != nil {
		return models.AuthResponse{}, err
	}
	defer resp.Body.Close()

	var result models.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
