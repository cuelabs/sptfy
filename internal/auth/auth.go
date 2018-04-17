package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
	"os"
	"os/exec"
	"runtime"
)

const (
	SPTFY_SERVER_ADDRESS string = "https://sptfy.cue.zone"
	SPTFY_CLIENT_ID      string = "940383534de04a41b61c51cbbd550708"
	SPTFY_REDIRECT_URI   string = "https://sptfy.cue.zone/callback"
)

type Authentication struct {
	cachePath string
	token     *oauth2.Token
}

// First look for the token on disk
// If it doesn't exist, then execute the authentication
func (a *Authentication) Token() (*oauth2.Token, error) {
	// Load the token from the cache it it does not exist
	if a.token == nil {
		if err := a.Load(""); err != nil {
			if err != nil {
				if err = a.Authenticate(); err != nil {
					return nil, err
				}
			}
		}
	}
	return a.token, nil
}

func (a *Authentication) Authenticate() error {
	// Load and create the Oauth2 configuration
	config, err := a.Config()
	if err != nil {
		return err
	}

	authUrl := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Println("In order to authenticate, use a browser to authorize the sptfy CLI")

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", authUrl).Start()

	case "windows", "darwin":
		err = exec.Command("open", authUrl).Start()
	default:
		err = fmt.Errorf("Unsupported platform: %s", runtime.GOOS)
	}

	if err != nil {
		fmt.Printf("Copy and paste the following link: \n%s\n\n", authUrl)
	}
	var code string

	fmt.Println("Enter the authorization code: ")
	fmt.Scanln(&code)

	if code == "" {
		return fmt.Errorf("Unable to read authorization code.")

	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return err
	}
	a.token = token
	a.Save()
	return nil
}

// This could be a potential source of really bad problems upcoming. Work on it often.
func (a *Authentication) Config() (*oauth2.Config, error) {
	scopes := []string{}
	scopes = append(scopes, "user-read-private", "streaming", "user-modify-playback-state")
	return &oauth2.Config{
		ClientID: SPTFY_CLIENT_ID,
		Scopes: scopes,
		Endpoint: spotify.Endpoint,
		RedirectURL: SPTFY_REDIRECT_URI,
		ClientSecret: os.Getenv("SPTFY_CLIENT_SECRET"),
	}, nil
}

func (a *Authentication) Cache(token *oauth2.Token) error {

}

// Returns an error if token cannot be loaded from cache.
func (a *Authentication) Load(path string) error {
	var err error
	if path == "" {
		path, err = a.CachePath()
		if err != nil {
			return err
		}
	} else {
		a.cachePath = path
	}

	// open the file at the path
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Could not open cache file at %s: %v", path, err)
	}
	defer f.Close()

	a.token = new(oauth2.Token)
	if err := json.NewDecoder(f).Decode(a.token); err != nil {
		return fmt.Errorf("Could not decode token in cache file at %s: %v", path, err)
	}
	return nil
}

func (a *Authentication) CachePath() (string, error) {
	if a.cachePath == "" {
		// look up home directory
	}
}

func (a *Authentication) ConfigPath() (string, error) {
	if a.configPath {

	}
}
