package main

import (
	"flag"
	"fmt"
	//"golang.org/x/crypto/ssh/terminal"
	// "github.com/cuelabs/sptfy/pkg/user"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	SPTFY_CLIENT_ID    string = "940383534de04a41b61c51cbbd550708"
	SPTFY_REDIRECT_URI string = "https://sptfy.cue.zone/callback"
	SPTFY_SCOPE_SET    string = "'user-read-private'%20'streaming'"
	SPTFY_STATE_PSK    string = "random"
)

type Ennvars struct {
	Version string
}

type Auth struct {
	cachePath string
	token     *oauth2.Token
}

type Environment struct {
	auth    Auth
	envvars Ennvars
	log     *log.Logger
}

func (a *Auth) Token(*oauth2.Token, error) {
	if a.token == nil {

	}
}

var env Environment

func init() {
	env.log = log.New(os.Stdout, "SPTFY", 0)
	initLog := env.log
	initLog.Println("Initiated SPTFY logging")
	authSpotifyUrl := url.URL{
		Scheme: "https",
		Host:   "accounts.spotify.com/authorize",
		Opaque: fmt.Sprintf("/?client_id=%v&response_type=code&redirect_uri=%v&state=%v&scopts=%v",
			SPTFY_CLIENT_ID,
			SPTFY_REDIRECT_URI,
			SPTFY_STATE_PSK,
			SPTFY_SCOPE_SET)}
	//req, err := http.NewRequest("GET", authSpotifyUrl.String(), nil)
	//if err != nil {
	//	initLog.Println("Unable to craft a Spotify API authorization request. Exiting")
	//	initLog.Print(err)
	//	os.Exit(1)
	//}
	resp, err := http.Get(authSpotifyUrl.String())
	if err != nil {
		initLog.Println("Error requesting Spotify authorization. Exiting")
		initLog.Print(err)
		os.Exit(1)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print()
	}
	fmt.Println(r)
}

func main() {

	info := flag.Bool("info", false, "See Spotify user information.")
	auth := flag.Bool("auth", false, "Authenticate your Spotify account.")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	playCmd := flag.NewFlagSet("play", flag.ExitOnError)

	searchType := searchCmd.String("type", "track", "Type of search <track|artist|albums>. Default: track.")
	searchQuery := searchCmd.String("query", "", "Search query. Use quotation marks if query contains spaces. (Required)")

	tag := playCmd.String("tag", "", "Playback item with sptfy tag. (Required)")

	if len(os.Args) < 2 {
		flag.Parse()
		if !(*info || *auth) {
			log.Println()
			flag.PrintDefaults()
			os.Exit(1)
		}
		if *info {
			// check if token exists

			// if not, authorize

			// hit the url of user data
			io.WriteString(os.Stdout, "IMPLEMENT")
		}
		flag.PrintDefaults()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "search":
		searchCmd.Parse(os.Args[:2])
	case "play":
		playCmd.Parse(os.Args[:2])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if searchCmd.Parsed() {
		if *searchQuery == "" {
			searchCmd.PrintDefaults()
			os.Exit(1)
		}
		searchUrl := url.URL{Scheme:
			}
	}

	if playCmd.Parsed() {
		if *tag == "" {
			flag.PrintDefaults()
		}
	}
}

/*
// Send CLI user login info to authorization server
//
func Authenticate(email string, password []byte) (token string, err error) {
	p := fmt.Sprintf("/?client_id=%v&response_type=token&scope=%v&show_dialog=false&redirect_uri=https://sptfy.cue.zone/redirect",
		SPTFY_CLIENT_ID,
		SPTFY_SCOPE_SET)
	//body := fmt.Sprintf("email=%s&password=%s", email, password)

	u := url.URL{Scheme: "https", Host: "accounts.spotify.com/authorize", Path: p}
	req, err := http.NewRequest("POST", u, nil)
	if err != nil {
		fmt.Println("")
	}
	resp, err := http.Client.Do(req)
	r, err := ioutil.ReadAll(resp.Body)
	fmt.Println(r)
	return "", nil
}*/

func makeAccessHeader(access_token string) http.Header {
	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("xi-li-format", "json")
	header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))
	return header
}
