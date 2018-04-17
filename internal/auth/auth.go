package auth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"os"
	"runtime"
	"os/exec"
)

type Authentication struct {
	cachePath string
	token     *oauth2.Token
}

func (a *Authentication) Token() (*oauth2.Token, error) {
	// load the token from the cache it it does not exist
	if a.token == nil {
		if err := a.Load(""); err != nil {

		}
	}
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

code, err :=
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

func (a *Authentication) Config() (*oauth2.Config, error) {
	path, err := a.ConfigPath()
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
