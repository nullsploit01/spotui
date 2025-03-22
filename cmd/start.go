package cmd

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

const (
	clientID    = "a33f9600ad424318ab1871106cc31761"
	redirectURI = "http://localhost:5555/callback"
	scopes      = "user-read-private user-read-email"
	state       = "randomstate"
)

func Start() {
	verifier, err := generateCodeVerifier()
	if err != nil {
		log.Fatal(err)
	}

	challenge := generateCodeChallenge(verifier)
	authURL := getAuthURL(clientID, redirectURI, challenge, state)

	fmt.Println("Opening browser for authentication...")
	exec.Command("open", authURL).Start() // For Linux; use "open" on macOS or "start" on Windows

	code, err := startRedirectServer()
	if err != nil {
		log.Fatal(err)
	}

	token, err := exchangeCodeForToken(code, verifier, clientID, redirectURI)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Access Token Response:")
	prettyPrintJSON(token)
}

func generateCodeVerifier() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func generateCodeChallenge(verifier string) string {
	h := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h[:])
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

func exchangeCodeForToken(code, verifier, clientID, redirectURI string) (map[string]interface{}, error) {
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

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func prettyPrintJSON(data interface{}) {
	out, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(out))
}
