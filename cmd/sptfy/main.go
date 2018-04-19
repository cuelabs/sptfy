package main

import (
	"flag"
	"fmt"
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/internal/spotifyclient"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	SPTFY_CLIENT_ID    string = "940383534de04a41b61c51cbbd550708"
	SPTFY_HOST         string = "sptfy.cue.zone"
	SPTFY_REDIRECT_URI string = "https://sptfy.cue.zone/callback"
	SPTFY_SCOPE_SET    string = "'user-read-private'%20'streaming'%20'user-modify-playback-state'"
	SPTFY_STATE_PSK    string = "random"
	SPTFY_CLIENT_TYPE  string = "SpotifyHttp" // "SptfyRpc"

)

var vars = &environment.Envvars{
	Version:    "0.0.1",
	ClientType: "SpotifyHttp",
}

var env *environment.Environment

func init() {
	// Environment variables gathered from environment
	env = &environment.Environment{
		nil,
		vars,
		log.New(os.Stdout, "SPTFY", 0),
		nil}
	env.Log.Println("Initiated SPTFY logging")
	authSpotifyUrl := url.URL{
		Scheme: "https",
		Host:   "accounts.spotify.com/authorize",
		Opaque: fmt.Sprintf("/?client_id=%v&response_type=code&redirect_uri=%v&state=%v&scopts=%v",
			SPTFY_CLIENT_ID,
			SPTFY_REDIRECT_URI,
			SPTFY_STATE_PSK,
			SPTFY_SCOPE_SET)}
	resp, err := http.Get(authSpotifyUrl.String())
	if err != nil {
		env.Log.Println("Error requesting Spotify authorization. Exiting.")
		env.Log.Print(err)
		os.Exit(1)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Could not request authorization")
	}

	// Give the environment a client from which to call Spotify
	s := fmt.Sprintf("%vClient", env.Vars.ClientType)
	switch {
	case s == "SpotifyHttpClient":
		env.Client = &spotifyclient.SpotifyHttpClient{}
	/*
		case s == "SptfyRpcClient":
		env.Client = &spotifyclient.SptfyRpcClient{&url.URL{Scheme: "https", Host: SPTFY_HOST}}
	*/
	default:
		env.Log.Println("Client type not valid. Check environment variables. Exiting.")
		fmt.Println("Client type not valid. Exiting.")
		os.Exit(1)
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

	// check for access_token at ~/.sptfy/token.json
	// if it exists

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
		switch *searchType {
		case "":
			a, err := env.Client.SearchArtist(*searchQuery, env)
			if err != nil {
				env.Log.Println("Failed SearchArtist()")
				fmt.Println()
			}

			// marshall to artists CLI output response

		}

	}

	if playCmd.Parsed() {
		if *tag == "" {
			flag.PrintDefaults()
		}
	}
}


func makeAccessHeader(access_token string) http.Header {
	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("xi-li-format", "json")
	header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))
	return header
}
