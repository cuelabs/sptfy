package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"net/http"
	"golang.org/x/crypto/ssh/terminal"
	"bytes"
)

func main() {
	info := flag.Bool("info", true, "See Spotify user information.")
	auth := flag.Bool("auth", true, "Authenticate your Spotify account.")
/*
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	playCmd := flag.NewFlagSet("play", flag.ExitOnError)

	searchType := searchCmd.String("type", "track", "Type of search <track|artist|albums>. Default: track.")
	searchQuery := searchCmd.String("query", "", "Search query. Use quotation marks if query contains spaces. (Required)")

	tag := playCmd.String("tag", "", "Playback item with sptfy tag. (Required)")
*/
	if len(os.Args) < 2 {
		flag.Parse()
		if !(*info || *auth) {
			flag.PrintDefaults()
			os.Exit(1)
		}
		if *info {
			// retrieve user's access token (check for existence of file)
			// request to
			io.WriteString(os.Stdout, "IMPLEMENT")
		}
		if *auth {
			// GET the Spotify auth URL
			var email string
			fmt.Println("Enter Spotify login email: ")
			_, err := fmt.Scanln(&email)
			if err != nil {
				fmt.Println("Unable to read email")
				fmt.Println(err)
			}
			fmt.Println("Enter Spotify login password: ")
			password, err := terminal.ReadPassword(0)
			if err != nil {
				fmt.Println("Unable to read password")
				fmt.Println(err)
				flag.PrintDefaults()
				os.Exit(1)
			}
            t, err := Authorization(email, password)
            if err != nil {
            	fmt.Println("Unable to authorize")
            	fmt.Println(err)
            	os.Exit(1)
			}
			// store token locally

		}
	}
}

// Send CLI user login info to authorization server
//
func Authorization(email string, password []byte) (token string, err error) {
    body := fmt.Sprintf("email=%s&password=%s", email, password)
	resp, err := http.Client{}.Post(
    	"https://sptfy.cue.zone/auth",
    	"application/json",
    	bytes.NewBuffer([]byte(body)))
	if err != nil {
        fmt.Println("")
	}
}

func makeAccessHeader(access_token string) http.Header {
	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("xi-li-format", "json")
	header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))
	return header
}

func authorize() (string, error) {

}
