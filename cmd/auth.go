package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	spotifyoauth2 "golang.org/x/oauth2/spotify"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

type Authentication struct {
	token     *oauth2.Token
	cachePath string
}

// Look in cache
// Authentication if not in cache

func (auth *Authentication) Token() (*oauth2.Token, error) {
	if auth.token == nil {
		if err := auth.Load(""); err != nil {
			if err = auth.Authenticate(); err != nil {
				return nil, err
			}
		}
		return nil, fmt.Errorf("Failed to load token or authenticate")
	}
	return auth.token, nil
}

func (auth *Authentication) Authenticate() error {
	config, err := auth.Config()
	if err != nil {
		return err
	}
	authUrl := config.AuthCodeURL(spotifyoauth2.Endpoint.AuthURL)

	fmt.Println("In order to authenticate, use a broswer to authorize sptfy with Spotify")

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", authUrl).Start()

	case "windows", "darwin":
		err = exec.Command("open", authUrl).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	if err != nil {
		fmt.Printf("Could not open the browser")
		fmt.Printf("copy and paste the following link \n%s\n\n", authUrl)

	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter authorization code:")
	code, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read authorization code", err)
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web", err)
	}

	auth.token = token
	return nil

}

func ConfigFromEnvironment() (*oauth2.Config, error) {
	return &oauth2.Config{
		ClientID:     *confSptfy.SpotifyClientId,
		ClientSecret: *confSptfy.SpotifyClientSecret,
		RedirectURL:  "https://sptfy.cue.zone/oauth2/callback",
		Endpoint: oauth2.Endpoint{spotifyoauth2.Endpoint.AuthURL,
			spotifyoauth2.Endpoint.TokenURL,
		},
	}, nil
}

// Config loads the client_secret.json from the ConfigPath.
func (auth *Authentication) Config() (*oauth2.Config, error) {
	return ConfigFromEnvironment()
}

func (auth *Authentication) CachePath() (string, error) {
	if auth.cachePath == "" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		cacheDir := filepath.Join(usr.HomeDir, ".sptfy")
		os.MkdirAll(cacheDir, 0700)

		cacheFile := url.QueryEscape("credentials.json")
		auth.cachePath = filepath.Join(cacheDir, cacheFile)
	}

	return auth.cachePath, nil
}

func (auth *Authentication) Load(path string) error {
	var err error
	if path == "" {
		path, err = auth.CachePath()
		if err != nil {
			return err
		}
	} else {
		auth.cachePath = path
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open cache file at %s: %v", path, err)
	}
	defer f.Close()

	auth.token = new(oauth2.Token)
	if err := json.NewDecoder(f).Decode(auth.token); err != nil {
		return fmt.Errorf("could not decode token in cache file at %s: %v")
	}
	return nil
}

func (auth *Authentication) Save(path string) error {
	var err error

	if path == "" {
		path, err = auth.CachePath()
		if err != nil {
			return err
		}
	} else {
		auth.cachePath = path
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Errorf("unable to catch oauth token: %v", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(auth.token); err != nil {
		return fmt.Errorf("could not encode oauth token: %v", err)
	}

	return nil
}

func (auth *Authentication) Delete(path string) {
	if path == ""{
		path, _ = auth.CachePath()
	} else {
		auth.cachePath = path
	}
	os.Remove(path)
}
