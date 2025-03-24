package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	"github.com/nullsploit01/spotui/cmd/utils"
)

var (
	scopes = "user-read-private user-read-email"
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

	authURL := getAuthURL(app.config.ClientID, app.config.RedirectURI, challenge, state)

	fmt.Println("Opening browser for authentication...")
	exec.Command("open", authURL).Start() // For Linux; use "open" on macOS or "start" on Windows

	code, err := startRedirectServer()
	if err != nil {
		return err
	}

	token, err := exchangeCodeForToken(code, verifier, app.config.ClientID, app.config.RedirectURI)
	if err != nil {
		return err
	}

	fmt.Println("Access Token Response:")
	prettyPrintJSON(token)
	return nil
}

func getAuthURL(clientID, redirectURI, challenge, state string) string {
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

func startRedirectServer() (string, error) {
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

func exchangeCodeForToken(code, verifier, clientID, redirectURI string) (map[string]any, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("code_verifier", verifier)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func prettyPrintJSON(data any) {
	out, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(out))
}
