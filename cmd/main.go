package cmd

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
	"context"
	"net/url"
	"net/http"
	"github.com/gogo/protobuf/io"
	"io/ioutil"
)

type Config struct {
	SpotifyClientId     *string
	SpotifyClientSecret *string
}

var confSptfy Config

func init() {
	// read in configuration variables from the environment
	if err := envconfig.Process("sptfy", &confSptfy); err != nil {
		fmt.Println("Could not process environment variables")
		os.Exit(1)
	}
	// configuration checks
	if (*confSptfy.SpotifyClientId == "" || *confSptfy.SpotifyClientSecret == "" ) {
		fmt.Println("SPTFY_SPOTIFYCLIENTID and SPTFY_SPOTIFYCLIENTSECRET must be set")
		os.Exit(1)
	}
	// authorize to Spotify
	auth := new(Authentication)
	config, err := auth.Config()
	if err != nil {
		fmt.Println("Could not configure Spotify authentication")
		os.Exit(1)
	}

	token, err := auth.Token()
	if err != nil {
		fmt.Println("Could not get Spotify authentication token")
		os.Exit(1)
	}
client := config.Client(context.Background(), token)
	test := &url.URL{Scheme: "https", Host: "api.spotify.com", Path: "/v1/me"}
	r, err := client.Get(test.String())
    fmt.Println(r)
    fmt.Println("Successfully made a request as Spotify client")
}

func main() {
	playCmd := flag.NewFlagSet("play", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)

	info := flag.Bool("info", false, "Display Spotify login and playback information.")

	playTrack := playCmd.String("track", "", "Play a track with sptfy tag")
	playAlbum := playCmd.String("album", "", "Play an album with sptfy tag")
	playArtist := playCmd.String("artist", "", "Play an artist with sptfy tag")
	playNext := playCmd.Bool("next", true, "Play next track")

	searchType := searchCmd.String("type", "tracks", "Type of search <tracks|artists|albums|playlists>. Default: tracks.")
    searchQuery := searchCmd.String("query", "", "Search query")

	if len(os.Args) < 2 {
		flag.Parse()
		if !(*info) {
			fmt.Println("Not a flag\n")
			flag.PrintDefaults()
			os.Exit(1)
		}
		fmt.Println("IMPLEMENT THE INFO COMMAND")
	}

	switch os.Args[1] {
	case "play":
		playCmd.Parse(os.Args[2:])
	case "search":
		searchCmd.Parse(os.Args[2:])
	default:
		fmt.Println("Not a flag\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if playCmd.Parsed() {
		if !(*playTrack == "" || *playAlbum == "" || *playArtist == "" || *playNext) {
           fmt.Println("No correct flags provided\n")
           flag.PrintDefaults()
           os.Exit(1)
		}
	}

	if searchCmd.Parsed() {
		if *searchType == "" {
			fmt.Println("Must include search type <tracks|artists|albums|playlists>\n")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
}
