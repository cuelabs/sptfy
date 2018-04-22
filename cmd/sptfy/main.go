package main

import (
	"flag"
	"fmt"
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/internal/spotifyclient"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/track"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	SPTFY_CLIENT_ID     string = "940383534de04a41b61c51cbbd550708"
	SPTFY_HOST          string = "sptfy.cue.zone"
	SPTFY_REDIRECT_URI  string = "https://sptfy.cue.zone/callback"
	SPTFY_STATE_PSK     string = "random"
	SPOTIFY_CLIENT_TYPE string = "SpotifyHTTP" // or SptfyRPC
	SPTFY_CACHE_PATH    string = "~/.sptfy/token.json"
)

var env *environment.Environment

func init() {
	// Variables configured in init()
	env = &environment.Environment{
		nil,
		nil,
		nil,
		nil,
		nil}
	env.Vars = &environment.Envvars{}

	// Logging for container environment
	env.Log = log.New(os.Stdout, "SPTFY", 0)
	env.Log.Println("Initiated SPTFY logging")

	// Set configuration settings in the application env
	env.Log.Println("Reading configuration settings")
	env.Settings = &environment.SptfySettings{}
	env.Settings.SetSetting("authorization-scope-set", []string{
		"user-read-private",
		"user-read-playback-state",
		"user-modify-playback-state",
		"streaming",
		"user-read-recently-played",
		"user-library-read"})
	env.Settings.SetSetting("sptfy-hosts", []string{SPTFY_HOST})
	env.Settings.SetSetting("sptfy-client-type", []string{})
	env.Log.Println("Finished setting settings in application env")

	// Authentication the user to Spotify API
	// User allows SPTFY Spotify API authorizations

	// Retrieve scopes from settings
	scopeset, ok := env.Settings.GetSetting("authorization-scope-set")
	if !ok {
		env.Log.Println("Authorization scopes not set in settings. Exiting")
		os.Exit(1)
	}
	scopesI := scopeset.([]interface{})
	scopeSlice := make([]string, len(scopesI))
	for _, v := range scopesI {
		scopeSlice = append(scopeSlice, v.(string))
	}
	scopeString := strings.Join(scopeSlice, "%20")

	// Send a request to the Spotify authorization service
	authSpotifyUrl := url.URL{
		Scheme: "https",
		Host:   "accounts.spotify.com/authorize",
		Opaque: fmt.Sprintf("/?client_id=%v&response_type=code&redirect_uri=%v&state=%v&scopes=%v",
			SPTFY_CLIENT_ID,
			SPTFY_REDIRECT_URI,
			SPTFY_STATE_PSK,
			scopeString)}
	resp, err := http.Get(authSpotifyUrl.String())
	// Open in broswer?
	if err != nil {
		env.Log.Println("Error requesting Spotify authorization. Exiting.")
		env.Log.Print(err)
		os.Exit(1)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Could not request authorization")
	}

	// Somehow needs to get initial access token

	// Set a Spotify client in the application env
	ct, ok := env.Settings.GetSetting("spotify-client-type")
	if !ok {
		ct = "SpotifyHTTP"
	}
	s := fmt.Sprintf("%vClient", ct)
	switch {
	case s == "SpotifyHttpClient":
		env.Client = &spotifyclient.SpotifyHttpClient{}
	/*
		    NextMonth 18/05
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
			inf, err := env.Client.RetrieveInfo(env)
			if err != nil {
				env.Log.Print(err)
				env.Log.Println("Unable to call env.Client.RetrieveInfo() while parsing flags in main()")
				fmt.Println("ERROR: Unable to get info")
			}

		}
		if *auth {
			if err := env.Client.RetrieveAuth(env); err != nil {
				env.Log.Print(err)
				env.Log.Println("Unable to call env.Client.RetrieveAuth() while parsing flags in main()")
				fmt.Println("ERROR: Unable to authorize")
			}
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
		case "", "track":

			// Search tracks with the SpotifyAPIOperations client
			tracks, err := env.Client.SearchTrack(*searchQuery, env)
			if err != nil {
				env.Log.Println("Failed SearchTrack() on searchCmd()")

			}

			// Iterate through the tracks and marshall them to a display format
			respHead := fmt.Sprintf("Displaying track search results for query %v\n", searchQuery)
			cliResp := &track.SptfyCLITrackDisplayResponse{
				respHead,
				"%v -- %v\t\t%v\n",
				[]*track.SptfyCLITrackDisplayItem{},
			}
			for _, t := range tracks {
				tag := &spotifyclient.SptfyTag{}
				if err := tag.New(); err != nil {
					env.Log.Print(err)
					env.Log.Println("Failed to generate a new tag. Exiting")
					os.Exit(1)
				}
				artists := t.Artists
				as := []*string{}
				for _, artist := range artists {
					as = append(as, artist.Name)
				}
				i := &track.SptfyCLITrackDisplayItem{
					tag,
					t.Name,
					as,
					t.Album.Name,
				}
				cliResp.Results = append(cliResp.Results, i)
			}

			// Print the result to output
			fmt.Println(respHead)
			for _, res := range cliResp.Results {
				fmt.Printf(cliResp.MessageFormat, res.Tag, res.TrackName, res.AlbumName)
			}
			os.Exit(0) // succcess!
			// END OF THIS FUNCTION

	    // Search albums with the SpotifyAPIOperations client
		case "album":
			albums, err := env.Client.SearchAlbum(*searchQuery, env)
			if err != nil {
				env.Log.Println("Failed SearchTrack() on searchCmd()")

			}

			respHead := fmt.Sprintf("Displaying album results for query %v", searchQuery)
			cliResp := &album.SptfyCLIAlbumDisplayResponse{
				respHead,
				"%v -- %v\t\t%v",
				[]*album.SptfyCLIAlbumDisplayItem{},
			}
			for _, a := range albums {
				tag := &spotifyclient.SptfyTag{}
				if err := tag.New(); err != nil {
					env.Log.Print(err)
					env.Log.Println("Failed to generate a new tag. Exiting")
					os.Exit(1)
				}
				artists := a.Artists
				as := []*string{}
				for _, artist := range artists {
					as = append(as, )
				}

			}

		}

		// Iterate through the albums and marshall them to a display format


		case "artist":
			artists, err := env.Client.Sea

		respHead := fmt.Sprintf("Displaying album results for query %v", searchQuery)
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
