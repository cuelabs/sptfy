package spotifyclient

import (
	"net/url"
	"github.com/cuelabs/sptfy/internal/auth"
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"github.com/cuelabs/sptfy/pkg/user"
	"net/http"
	"fmt"
	"context"
	"io/ioutil"
	"errors"
)

const (
	SPTFY_CACHE_PATH string = "~/.sptfy/token.json"
)

type SpotifyApiOperations interface {
	RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error)

	RetrieveAuth(e *environment.Environment) error

	PlaybackNext(e *environment.Environment) (*track.SptfyTrack, error)
	PlaybackPlay(e *environment.Environment) (*track.SptfyTrack, error)
	PlaybackPause(e *environment.Environment) (*track.SptfyTrack, error)

	// Do I need to pass in auth.Authentication for this?
	SearchAlbum(query string, e *environment.Environment) (*album.SptfyAlbum, error)
	SearchArtist(query string, e *environment.Environment) (*artist.SptfyArtist, error)
	SearchTrack(query string, e *environment.Environment) (*track.SptfyTrack, error)
}

type SpotifyHttpClient struct {}

type SptfyRpcClient struct {
	SptfyHost *url.URL
}

func (s *SpotifyHttpClient) RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error) {
	// Create a client
	token, err := e.Auth.Token()
	 if err != nil {
	 	e.Log.Println("Unable to retrieve token in RetrieveInfo()")
	 	return nil, err
	 }
	 conf, err := e.Auth.Config()
	 if err != nil {
	 	e.Log.Println("Unable to create retrieve OAuth2 config in RetrieveInfo()")
	 	return nil, err
	 }
	 client := conf.Client(context.Background(), token)
	u := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: "/v1/me",
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		e.Log.Println("Unable to create a new request in RetrieveInfo()")
		return nil, err
	}
	 resp, err := client.Do(req)
	 if err != nil {
	 	e.Log.Println("Request failed in RetrieveInfo()")
	 }
	 r, err := ioutil.ReadAll(resp.Body)
	 // marshall data response to


	 fmt.Printf("This is the response: %v", r)

	 return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) RetrieveAuth(e *environment.Environment) error {
	if e.Auth == nil {
        e.Log.Println("Authorization not found. Beginning authentication.")
        a := &auth.Authentication{SPTFY_CACHE_PATH, nil}
		fmt.Print("Please authenticate")
        a.Authenticate()
        e.Log.Println("Authentication complete.")
        fmt.Println("Thank you for authenticating.")
	}

	return errors.New("Not implemented")
}

// These handlers control playback with a Spotify
func (s *SpotifyHttpClient) PlaybackNext(e *environment.Environment) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPlay(e * environment.Environment) (*track.SptfyTrack, error) {
    return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPause(e *environment.Environment) (*track.SptfyTrack, error) {
    return nil, errors.New("Not implemented")
}

// These handlers search the Spotify API with(out) authentication
func (s *SpotifyHttpClient) SearchAlbum(query string, e *environment.Environment) (*album.SptfyAlbum, error) {
	// search path
	sp := fmt.Sprintf("/v1/search?q=%v&type=album", query)
	// search url
	su := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: sp}
		resp, err := http.Get(su.String())
		if err != nil {
			return nil, err
			}
			defer resp.Body.Close()
			// no auth?
			//
    b, err := ioutil.ReadAll(resp.Body)
    fmt.Println(b)
    fmt.Println("Got to the end of SearchAlbum()")
    return nil, errors.New("Implemented but not tested")
}

func (s *SpotifyHttpClient) SearchArtist(query string, e *environment.Environment) (*artist.SptfyArtist, error) {
	// sanitize query string; replace spaces with '%20'
	sp := fmt.Sprintf("/v1/search?q=%v&type=artist", query)
    // search url
    u := &url.URL{
    	Scheme: "https",
    	Host: "api.spotify.com",
    	Path: sp,
	}
	resp, err := http.Get(u.String())
	if err != nil {
        return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(b)
	fmt.Println("Got to end of SearchArtistr()")
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) SearchTrack(query string, e *environment.Environment) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented.")
}
